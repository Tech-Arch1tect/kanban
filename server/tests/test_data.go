package tests

import "server/models"

var (
	TestAdminUser = &models.User{
		Username: "test_admin",
		Email:    "admin@example.local",
		Password: "password123",
		Role:     models.RoleAdmin,
	}
	TestUser = &models.User{
		Username: "test_user",
		Email:    "user@example.local",
		Password: "password123",
		Role:     models.RoleUser,
	}
	TestBoard = &models.Board{
		Name: "Test Board",
		Slug: "test-board",
	}
)
