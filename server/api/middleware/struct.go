package middleware

import (
	"server/database/repository"
	"server/internal/helpers"
)

type Middleware struct {
	db     *repository.Database
	helper *helpers.HelperService
}

func NewMiddleware(db *repository.Database, helper *helpers.HelperService) *Middleware {
	return &Middleware{db: db, helper: helper}
}
