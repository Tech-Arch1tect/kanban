package board

import (
	"server/database/repository"
	"server/models"
	"server/services/role"
	"sort"
)

func (bs *BoardService) GetBoard(id uint) (models.Board, error) {
	board, err := bs.db.BoardRepository.GetByID(id, repository.WithPreload("Swimlanes"), repository.WithPreload("Columns"))
	if err != nil {
		return models.Board{}, err
	}

	bs.SortBoard(&board)
	return board, nil
}

func (bs *BoardService) GetBoardBySlug(slug string) (models.Board, error) {
	board, err := bs.db.BoardRepository.GetFirst(repository.WithWhere("slug = ?", slug), repository.WithPreload("Swimlanes"), repository.WithPreload("Columns"))
	if err != nil {
		return models.Board{}, err
	}

	bs.SortBoard(&board)
	return board, nil
}

func (bs *BoardService) SortBoard(board *models.Board) {
	sort.Slice(board.Swimlanes, func(i, j int) bool { return board.Swimlanes[i].Order < board.Swimlanes[j].Order })
	sort.Slice(board.Columns, func(i, j int) bool { return board.Columns[i].Order < board.Columns[j].Order })
}

func (bs *BoardService) ListBoards(userID uint) ([]models.Board, error) {
	user, err := bs.db.UserRepository.GetByID(userID)
	if err != nil {
		return nil, err
	}

	if user.Role == models.RoleAdmin {
		return bs.db.BoardRepository.GetAll(repository.WithPreload("Swimlanes", "Columns"))
	}

	MemberRole, err := bs.db.BoardRoleRepository.GetFirst(repository.WithWhere("name = ?", role.MemberRole.Name))
	if err != nil {
		return nil, err
	}

	userBoardRoles, err := bs.db.UserBoardRoleRepository.GetAll(repository.WithWhere("user_id = ? AND board_role_id = ?", userID, MemberRole.ID), repository.WithPreload("Board", "Board.Swimlanes", "Board.Columns"))
	if err != nil {
		return nil, err
	}

	boards := make([]models.Board, len(userBoardRoles))
	for i, userBoardRole := range userBoardRoles {
		boards[i] = userBoardRole.Board
	}

	return boards, nil
}
