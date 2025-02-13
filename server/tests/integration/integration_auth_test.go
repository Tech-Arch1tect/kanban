package integration

import (
	"log"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

// this test registers a new user, logs in, and retrieves the profile
func TestRegisterLoginProfile(t *testing.T) {
	router := getTestingRouter()
	var cookies []*http.Cookie

	registrationData := map[string]string{
		"username": "integration_user",
		"email":    "integration@example.com",
		"password": "password123",
	}
	_, cookies = doAPIRequest(t, router, "POST", "/api/v1/auth/register", registrationData, cookies, http.StatusCreated)

	loginData := map[string]string{
		"email":    "integration@example.com",
		"password": "password123",
	}
	_, cookies = doAPIRequest(t, router, "POST", "/api/v1/auth/login", loginData, cookies, http.StatusOK)

	profile, _ := doAPIRequest(t, router, "GET", "/api/v1/auth/profile", nil, cookies, http.StatusOK)

	log.Println("Profile:", profile)

	require.Equal(t, "integration_user", profile["username"], "username should match")
	require.Equal(t, "integration@example.com", profile["email"], "email should match")
}
