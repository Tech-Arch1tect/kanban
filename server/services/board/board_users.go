package board

import (
	"errors"
	"server/database/repository"
	"server/services/role"

	"gorm.io/gorm"
)

func (bs *BoardService) AddOrInviteUserToBoard(authUserID, boardID uint, email string, r role.AppRole) error {
	authUser, err := bs.db.UserRepository.GetByID(authUserID)
	if err != nil {
		return err
	}

	can, err := bs.rs.CheckRole(authUser.ID, boardID, role.AdminRole)
	if err != nil || !can {
		return errors.New("forbidden")
	}

	user, err := bs.db.UserRepository.GetFirst(repository.WithWhere("email = ?", email))
	if err == gorm.ErrRecordNotFound {
		return bs.InviteUserToBoard(boardID, email, r.Name)
	}

	return bs.AddUserToBoard(boardID, user.ID, r)
}

func (bs *BoardService) RemoveUserFromBoard(currentUserID, requestedUserID, boardID uint) error {
	can, err := bs.rs.CheckRole(currentUserID, boardID, role.AdminRole)
	if err != nil || !can {
		return errors.New("forbidden")
	}

	return bs.rs.RemoveRole(requestedUserID, boardID)
}
