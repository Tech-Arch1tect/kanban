package controllers

import "server/services"

var adminService services.AdminService
var authService services.AuthService

func Init() error {
	adminService = services.AdminService{}
	authService = services.AuthService{}
	return nil
}
