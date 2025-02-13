package integration

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"server/cmd/initHelper"
	"server/tests"
	"testing"

	"github.com/stretchr/testify/require"
)

func getCSRFToken(t *testing.T, router http.Handler) (string, []*http.Cookie) {
	req := httptest.NewRequest("GET", "/api/v1/auth/csrf-token", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	require.Equal(t, http.StatusOK, resp.Code, "failed to retrieve CSRF token")

	var result map[string]string
	err := json.Unmarshal(resp.Body.Bytes(), &result)
	require.NoError(t, err, "failed to parse CSRF token response")
	csrfToken, ok := result["csrf_token"]
	require.True(t, ok, "csrf_token not found in response")

	cookies := resp.Result().Cookies()
	require.NotEmpty(t, cookies, "no cookies returned with CSRF token")
	return csrfToken, cookies
}

func mergeCookies(oldCookies, newCookies []*http.Cookie) []*http.Cookie {
	cookieMap := make(map[string]*http.Cookie)
	for _, c := range oldCookies {
		cookieMap[c.Name] = c
	}
	for _, c := range newCookies {
		cookieMap[c.Name] = c
	}
	merged := make([]*http.Cookie, 0, len(cookieMap))
	for _, c := range cookieMap {
		merged = append(merged, c)
	}
	return merged
}

func doAPIRequest(t *testing.T, router http.Handler, method, endpoint string, requestData interface{}, cookies []*http.Cookie, expectedStatus int) (map[string]interface{}, []*http.Cookie) {
	var csrfToken string
	if method == "POST" || method == "PUT" || method == "DELETE" {
		var csrfCookies []*http.Cookie
		csrfToken, csrfCookies = getCSRFToken(t, router)
		cookies = mergeCookies(cookies, csrfCookies)
	}

	var requestBody []byte
	if requestData != nil {
		var err error
		requestBody, err = json.Marshal(requestData)
		require.NoError(t, err, "failed to marshal request data")
	}

	req := httptest.NewRequest(method, endpoint, bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")
	if csrfToken != "" {
		req.Header.Set("X-CSRF-Token", csrfToken)
	}

	for _, cookie := range cookies {
		req.AddCookie(cookie)
	}

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	require.Equal(t, expectedStatus, resp.Code, "%s %s should return status %d", method, endpoint, expectedStatus)

	respCookies := resp.Result().Cookies()
	cookies = mergeCookies(cookies, respCookies)

	var response map[string]interface{}
	if len(resp.Body.Bytes()) > 0 {
		err := json.Unmarshal(resp.Body.Bytes(), &response)
		require.NoError(t, err, "failed to parse response body")
	} else {
		response = make(map[string]interface{})
	}

	return response, cookies
}

func getTestingRouter() http.Handler {
	os.Setenv("SQLITE_FILE_PATH", "test.db")
	os.Setenv("InsertTestData", "true")
	router, _, cleanup := initHelper.SetupRouter()
	defer cleanup()
	return router
}

func Login(t *testing.T, r http.Handler, role string) []*http.Cookie {
	var email string
	password := "password123"
	if role == "admin" {
		email = tests.TestAdminUser.Email
	} else {
		email = tests.TestUser.Email
	}
	var cookies []*http.Cookie

	loginData := map[string]string{
		"email":    email,
		"password": password,
	}

	_, cookies = doAPIRequest(t, r, "POST", "/api/v1/auth/login", loginData, cookies, http.StatusOK)

	return cookies
}
