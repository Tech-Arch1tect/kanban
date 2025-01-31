package board

import (
	"errors"
	"server/services/role"
)

func (bs *BoardService) ChangeBoardRole(currentUserID, RequestedUserID, boardID uint, r string) error {
	can, err := bs.rs.CheckRole(currentUserID, boardID, role.AdminRole)
	if err != nil {
		return err
	}

	if !can {
		return errors.New("forbidden")
	}

	AppRole := role.AppRole{
		Name: r,
	}

	err = bs.rs.AssignRole(RequestedUserID, boardID, AppRole)
	if err != nil {
		return err
	}

	return nil
}
