package board

import (
	"errors"
	"server/models"
	"server/services/role"
)

func (bs *BoardService) RenameBoard(userID, boardID uint, name string) (models.Board, error) {
	board, err := bs.db.BoardRepository.GetByID(boardID)
	if err != nil {
		return models.Board{}, err
	}

	can, err := bs.rs.CheckRole(userID, board.ID, role.AdminRole)
	if err != nil || !can {
		return models.Board{}, errors.New("forbidden")
	}

	board.Name = name

	return board, bs.db.BoardRepository.Update(&board)
}

func (bs *BoardService) UpdateBoardSlug(userID, boardID uint, slug string) (models.Board, error) {
	board, err := bs.db.BoardRepository.GetByID(boardID)
	if err != nil {
		return models.Board{}, err
	}

	can, err := bs.rs.CheckRole(userID, board.ID, role.AdminRole)
	if err != nil || !can {
		return models.Board{}, errors.New("forbidden")
	}

	board.Slug = slug

	return board, bs.db.BoardRepository.Update(&board)
}
