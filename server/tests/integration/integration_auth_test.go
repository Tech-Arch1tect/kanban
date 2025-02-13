package integration

import (
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

// this test registers a new user, logs in, and retrieves the profile
func TestRegisterLoginProfile(t *testing.T) {
	router := getTestingRouter()

	csrfToken, cookies := getCSRFToken(t, router)
	username := "integration_user"
	email := "integration@example.com"
	password := "password123"

	registerUser(t, router, csrfToken, cookies, username, email, password)
	updatedCookies, _ := loginUser(t, router, csrfToken, cookies, email, password)
	profile := getProfile(t, router, updatedCookies)

	log.Println("Profile:", profile)

	require.Equal(t, username, profile["username"], "username should match")
	require.Equal(t, email, profile["email"], "email should match")
}
