package integration

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/http/httptest"
	"os"
	"server/cmd/initHelper"
	"server/tests"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

type APITestClient struct {
	client   *http.Client
	baseURL  string
	cleanup  func()
	closeSrv func()
}

func NewAPITestClient() (*APITestClient, error) {
	// Set up environment variables for test environment.
	os.Setenv("SQLITE_FILE_PATH", "test.db")
	os.Setenv("InsertTestData", "true")
	os.Setenv("COOKIE_SECURE", "false")
	os.Setenv("RATE_LIMIT_ENABLED", "false")
	os.Setenv("APP_ENVIRONMENT", "testing")

	router, _, cleanup := initHelper.SetupRouter()

	ts := httptest.NewServer(router)

	jar, err := cookiejar.New(nil)
	if err != nil {
		cleanup()
		ts.Close()
		return nil, err
	}

	client := &http.Client{
		Jar: jar,
	}

	return &APITestClient{
		client:   client,
		baseURL:  ts.URL,
		cleanup:  cleanup,
		closeSrv: ts.Close,
	}, nil
}

func (c *APITestClient) DoRequest(method, path string, body io.Reader, headers map[string]string) (*http.Response, error) {
	fullURL := c.baseURL + path
	if headers == nil {
		headers = make(map[string]string)
	}

	if method == http.MethodPost || method == http.MethodPut || method == http.MethodPatch || method == http.MethodDelete {
		token, err := c.GetCSRFToken()
		if err != nil {
			return nil, err
		}
		headers["X-CSRF-Token"] = token
	}

	if headers["Content-Type"] == "" {
		headers["Content-Type"] = "application/json"
	}

	req, err := http.NewRequest(method, fullURL, body)
	if err != nil {
		return nil, err
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	return c.client.Do(req)
}

func (c *APITestClient) GetCSRFToken() (string, error) {
	res, err := c.client.Get(c.baseURL + "/api/v1/auth/csrf-token")
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	var data struct {
		CSRFToken string `json:"csrf_token"`
	}
	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		return "", err
	}
	token := strings.TrimSpace(data.CSRFToken)
	if token == "" {
		return "", errors.New("empty CSRF token received")
	}
	return token, nil
}

func (c *APITestClient) Close() {
	if c.cleanup != nil {
		c.cleanup()
	}
	if c.closeSrv != nil {
		c.closeSrv()
	}
}

// Login helper function.
func Login(client *APITestClient, role string) error {
	var email string
	password := "password123"

	if role == "admin" {
		email = tests.TestAdminUser.Email
	} else {
		email = tests.TestUser.Email
	}

	loginData := map[string]string{
		"email":    email,
		"password": password,
	}
	payload, err := json.Marshal(loginData)
	if err != nil {
		return err
	}

	res, err := client.DoRequest(http.MethodPost, "/api/v1/auth/login", bytes.NewReader(payload), nil)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return errors.New("login failed: unexpected status code")
	}
	return nil
}

func decodeAndCloseResponseBody(t *testing.T, res *http.Response, v interface{}) {
	defer res.Body.Close()
	err := json.NewDecoder(res.Body).Decode(v)
	require.NoError(t, err)
}
