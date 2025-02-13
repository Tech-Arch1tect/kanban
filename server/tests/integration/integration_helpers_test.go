package integration

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"server/cmd/initHelper"
	"testing"
	"time"

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

func registerUser(t *testing.T, router http.Handler, csrfToken string, cookies []*http.Cookie, username, email, password string) {
	data := map[string]string{
		"username": username,
		"email":    email,
		"password": password,
	}
	body, err := json.Marshal(data)
	require.NoError(t, err)

	req := httptest.NewRequest("POST", "/api/v1/auth/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-CSRF-Token", csrfToken)
	for _, cookie := range cookies {
		req.AddCookie(cookie)
	}
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	require.Equal(t, http.StatusCreated, resp.Code, "registration should succeed")
}

func findCookieByName(cookies []*http.Cookie, name string) *http.Cookie {
	for _, cookie := range cookies {
		if cookie.Name == name {
			return cookie
		}
	}
	return nil
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

func loginUser(t *testing.T, router http.Handler, csrfToken string, cookies []*http.Cookie, email, password string) ([]*http.Cookie, *http.Cookie) {
	data := map[string]string{
		"email":    email,
		"password": password,
	}
	body, err := json.Marshal(data)
	require.NoError(t, err)

	req := httptest.NewRequest("POST", "/api/v1/auth/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-CSRF-Token", csrfToken)
	for _, cookie := range cookies {
		req.AddCookie(cookie)
	}
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	require.Equal(t, http.StatusOK, resp.Code, "login should succeed")

	loginCookies := resp.Result().Cookies()
	cookies = mergeCookies(cookies, loginCookies)
	sessionCookie := findCookieByName(cookies, "mysession")
	require.NotNil(t, sessionCookie, "session cookie not found")
	return cookies, sessionCookie
}

func getProfile(t *testing.T, router http.Handler, cookies []*http.Cookie) map[string]interface{} {
	req := httptest.NewRequest("GET", "/api/v1/auth/profile", nil)
	req.URL.Scheme = "https"
	for _, cookie := range cookies {
		req.AddCookie(cookie)
	}
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	require.Equal(t, http.StatusOK, resp.Code, "profile endpoint should succeed")

	var profile map[string]interface{}
	err := json.Unmarshal(resp.Body.Bytes(), &profile)
	require.NoError(t, err)
	return profile
}

func getTestingRouter() http.Handler {
	dateString := time.Now().Format("2006-01-02 15:04:05")
	os.Setenv("SQLITE_FILE_PATH", "test-"+dateString+".db")
	os.Setenv("InsertTestData", "true")
	router, _, cleanup := initHelper.SetupRouter()
	defer cleanup()
	return router
}
