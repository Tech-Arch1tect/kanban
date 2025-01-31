package board

import (
	"errors"
	"server/services/role"
)

func (bs *BoardService) GetUsersWithAccess(userID, boardID uint) ([]role.UserWithAppRole, error) {
	can, err := bs.rs.CheckRole(userID, boardID, role.MemberRole)
	if err != nil {
		return nil, err
	}
	if !can {
		return nil, errors.New("forbidden")
	}

	users, err := bs.rs.GetUsersWithAccessToBoard(boardID)
	if err != nil {
		return nil, err
	}

	return users, nil
}
