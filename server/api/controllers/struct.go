package controllers

import (
	"server/api/controllers/admin"
	"server/api/controllers/auth"
	"server/api/controllers/misc"
	"server/config"
	"server/database/repository"
	"server/internal/helpers"
	"server/services"
)

type Controllers struct {
	AuthController  *auth.AuthController
	AdminController *admin.AdminController
	MiscController  *misc.MiscController
}

func NewControllers(cfg *config.Config, authService *services.AuthService, adminService *services.AdminService, db *repository.Database, helperService *helpers.HelperService) *Controllers {
	return &Controllers{
		AuthController:  auth.NewAuthController(authService, db, helperService),
		AdminController: admin.NewAdminController(adminService),
		MiscController:  misc.NewMiscController(cfg),
	}
}
