package board

import "server/services/role"

func (bs *BoardService) AddUserToBoard(boardID, userID uint, r role.AppRole) error {
	return bs.rs.AssignRole(userID, boardID, r)
}
