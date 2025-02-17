package integration

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestAdminBoard tests that an admin can create/get/delete a board.
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

	res, err = client.DoRequest("GET", "/api/v1/boards/get/"+fmt.Sprintf("%v", board["id"].(float64)), nil, nil)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, res.StatusCode)

	res, err = client.DoRequest("GET", "/api/v1/boards/get-by-slug/"+board["slug"].(string), nil, nil)
	require.NoError(t, err)
	err = json.NewDecoder(res.Body).Decode(&response)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, res.StatusCode)
	require.Equal(t, "Test Board test", response["board"].(map[string]interface{})["name"], "board name should match")

	deleteBoardData := map[string]interface{}{
		"id": board["id"].(float64),
	}
	payload, err = json.Marshal(deleteBoardData)
	require.NoError(t, err)

	res, err = client.DoRequest("GET", "/api/v1/boards/list", nil, nil)
	require.NoError(t, err)
	err = json.NewDecoder(res.Body).Decode(&response)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, res.StatusCode)
	require.Len(t, response["boards"], 2)
	defer res.Body.Close()

	res, err = client.DoRequest("POST", "/api/v1/boards/delete", bytes.NewReader(payload), nil)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, res.StatusCode)
}

// TestUserBoard tests that a user can NOT create/get/delete a board.
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

	res, err = client.DoRequest("GET", "/api/v1/boards/get/1", nil, nil)
	require.NoError(t, err)
	require.Equal(t, http.StatusForbidden, res.StatusCode)

	res, err = client.DoRequest("GET", "/api/v1/boards/get-by-slug/test-board", nil, nil)
	require.NoError(t, err)
	require.Equal(t, http.StatusForbidden, res.StatusCode)

	res, err = client.DoRequest("GET", "/api/v1/boards/list", nil, nil)
	require.NoError(t, err)
	err = json.NewDecoder(res.Body).Decode(&response)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, res.StatusCode)
	require.Len(t, response["boards"], 0)
	defer res.Body.Close()

	deleteBoardData := map[string]interface{}{
		"id": 1,
	}
	payload, err = json.Marshal(deleteBoardData)
	require.NoError(t, err)

	res, err = client.DoRequest("POST", "/api/v1/boards/delete", bytes.NewReader(payload), nil)
	require.NoError(t, err)
	require.Equal(t, http.StatusUnauthorized, res.StatusCode)
}
