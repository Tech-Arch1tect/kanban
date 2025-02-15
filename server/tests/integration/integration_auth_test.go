package integration

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"server/tests"

	"github.com/stretchr/testify/require"
)

// TestRegisterLoginProfile registers a new user, logs in, and retrieves the profile.
func TestRegisterLoginProfile(t *testing.T) {
	client, err := NewAPITestClient()
	require.NoError(t, err)
	defer client.Close()

	registrationData := map[string]string{
		"username": "integration_user",
		"email":    "integration@example.com",
		"password": "password123",
	}
	payload, err := json.Marshal(registrationData)
	require.NoError(t, err)

	res, err := client.DoRequest("POST", "/api/v1/auth/register", bytes.NewReader(payload), nil)
	require.NoError(t, err)
	require.Equal(t, http.StatusCreated, res.StatusCode)
	res.Body.Close()

	loginData := map[string]string{
		"email":    "integration@example.com",
		"password": "password123",
	}
	payload, err = json.Marshal(loginData)
	require.NoError(t, err)

	res, err = client.DoRequest("POST", "/api/v1/auth/login", bytes.NewReader(payload), nil)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, res.StatusCode)
	res.Body.Close()

	res, err = client.DoRequest("GET", "/api/v1/auth/profile", nil, nil)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, res.StatusCode)
	defer res.Body.Close()

	var profile map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&profile)
	require.NoError(t, err)

	require.Equal(t, "integration_user", profile["username"], "username should match")
	require.Equal(t, "integration@example.com", profile["email"], "email should match")
}

// TestLoginUser logs in as a user and confirms the role is not admin.
func TestLoginUser(t *testing.T) {
	client, err := NewAPITestClient()
	require.NoError(t, err)
	defer client.Close()

	require.NoError(t, Login(client, "user"))

	res, err := client.DoRequest("GET", "/api/v1/auth/profile", nil, nil)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, res.StatusCode)
	defer res.Body.Close()

	var response map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&response)
	require.NoError(t, err)

	require.Equal(t, tests.TestUser.Email, response["email"], "email should match")
	require.Equal(t, tests.TestUser.Username, response["username"], "username should match")
	require.Equal(t, "user", response["role"], "role should match")
}

// TestLoginAdmin logs in as an admin and confirms the role is admin.
func TestLoginAdmin(t *testing.T) {
	client, err := NewAPITestClient()
	require.NoError(t, err)
	defer client.Close()

	require.NoError(t, Login(client, "admin"))

	res, err := client.DoRequest("GET", "/api/v1/auth/profile", nil, nil)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, res.StatusCode)
	defer res.Body.Close()

	var response map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&response)
	require.NoError(t, err)

	require.Equal(t, tests.TestAdminUser.Email, response["email"], "email should match")
	require.Equal(t, tests.TestAdminUser.Username, response["username"], "username should match")
	require.Equal(t, "admin", response["role"], "role should match")
}
