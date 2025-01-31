package board

import (
	"errors"
	"server/models"
	"server/services/role"
)

func (bs *BoardService) GetBoardWithPermissions(userID, boardID uint) (models.Board, error) {
	can, err := bs.rs.CheckRole(userID, boardID, role.MemberRole)
	if err != nil || !can {
		return models.Board{}, errors.New("forbidden")
	}

	return bs.GetBoard(boardID)
}

func (bs *BoardService) GetBoardBySlugWithPermissions(userID uint, slug string) (models.Board, error) {
	board, err := bs.GetBoardBySlug(slug)
	if err != nil {
		return models.Board{}, err
	}

	can, err := bs.rs.CheckRole(userID, board.ID, role.MemberRole)
	if err != nil || !can {
		return models.Board{}, errors.New("forbidden")
	}

	return board, nil
}
