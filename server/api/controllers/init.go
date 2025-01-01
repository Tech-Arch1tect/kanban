package controllers

import "server/services"

var adminService services.AdminService

func Init() error {
	adminService = services.AdminService{}
	return nil
}

