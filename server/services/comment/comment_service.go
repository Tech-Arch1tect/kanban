package comment

import (
	"server/database/repository"
	"server/internal/helpers"
	"server/services/role"
)

type CommentService struct {
	db *repository.Database
	rs *role.RoleService
	hs *helpers.HelperService
}

func NewCommentService(db *repository.Database, rs *role.RoleService, hs *helpers.HelperService) *CommentService {
	return &CommentService{db: db, rs: rs, hs: hs}
}
