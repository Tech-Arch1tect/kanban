package controllers

import "server/services"

var adminService services.AdminService
var authService services.AuthService

type Controllers struct {
	AdminController *AdminController
	AuthController  *AuthController
	MiscController  *MiscController
}

func Init() (cr *Controllers, err error) {
	adminService = services.AdminService{}
	authService = services.AuthService{}

	return &Controllers{
		AdminController: &AdminController{},
		AuthController:  &AuthController{},
		MiscController:  &MiscController{},
	}, nil
}
