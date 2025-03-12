package board

import (
	"server/services/role"
)

func (bs *BoardService) CanAdministrateBoard(userID uint, boardID uint) bool {
	can, err := bs.rs.CheckRole(userID, boardID, role.AdminRole)
	if err != nil {
		return false
	}

	return can
}
