package integration

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestAdminBoard tests that an admin can create/delete a board.
func TestAdminBoard(t *testing.T) {
	client, err := NewAPITestClient()
	require.NoError(t, err)
	defer client.Close()

	require.NoError(t, Login(client, "admin"))

	boardData := map[string]interface{}{
		"name":      "Test Board test",
		"slug":      "test-Board-test",
		"swimlanes": []string{"Backlog", "In Progress", "Done"},
		"columns":   []string{"To Do", "Doing", "Done"},
	}
	payload, err := json.Marshal(boardData)
	require.NoError(t, err)

	res, err := client.DoRequest("POST", "/api/v1/boards/create", bytes.NewReader(payload), nil)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, res.StatusCode)

	var response map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&response)
	require.NoError(t, err)

	board, ok := response["board"].(map[string]interface{})
	require.True(t, ok, "board response missing")
	require.Equal(t, "Test Board test", board["name"], "board name should match")
	require.Equal(t, "test-board-test", board["slug"], "board slug should be lowercased")

	deleteBoardData := map[string]interface{}{
		"id": board["id"].(float64),
	}
	payload, err = json.Marshal(deleteBoardData)
	require.NoError(t, err)

	res, err = client.DoRequest("POST", "/api/v1/boards/delete", bytes.NewReader(payload), nil)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, res.StatusCode)
}

// TestUserBoard tests that a user can NOT create/delete a board.
func TestUserBoard(t *testing.T) {
	client, err := NewAPITestClient()
	require.NoError(t, err)
	defer client.Close()

	require.NoError(t, Login(client, "user"))

	boardData := map[string]interface{}{
		"name":      "Test Board test user",
		"slug":      "test-board-test-user",
		"swimlanes": []string{"Backlog", "In Progress", "Done"},
		"columns":   []string{"To Do", "Doing", "Done"},
	}
	payload, err := json.Marshal(boardData)
	require.NoError(t, err)

	res, err := client.DoRequest("POST", "/api/v1/boards/create", bytes.NewReader(payload), nil)
	require.NoError(t, err)
	require.Equal(t, http.StatusUnauthorized, res.StatusCode)

	var response map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&response)
	require.NoError(t, err)

	deleteBoardData := map[string]interface{}{
		"id": 1,
	}
	payload, err = json.Marshal(deleteBoardData)
	require.NoError(t, err)

	res, err = client.DoRequest("POST", "/api/v1/boards/delete", bytes.NewReader(payload), nil)
	require.NoError(t, err)
	require.Equal(t, http.StatusUnauthorized, res.StatusCode)
}
