package services

import (
	"errors"
	"server/database/repository"
	"server/models"
	"sort"
)

type BoardService struct {
	db *repository.Database
	rs *RoleService
}

func NewBoardService(db *repository.Database, rs *RoleService) *BoardService {
	return &BoardService{
		db: db,
		rs: rs,
	}
}

func (bs *BoardService) CreateBoard(name, slug string, swimlaneNames, columnNames []string) (models.Board, error) {
	swimlanes := make([]models.Swimlane, len(swimlaneNames))
	for i, name := range swimlaneNames {
		swimlanes[i] = models.Swimlane{Name: name, Order: i}
	}

	columns := make([]models.Column, len(columnNames))
	for i, name := range columnNames {
		columns[i] = models.Column{Name: name, Order: i}
	}

	board := models.Board{
		Name:      name,
		Slug:      slug,
		Swimlanes: swimlanes,
		Columns:   columns,
	}

	err := bs.db.BoardRepository.Create(&board)
	if err != nil {
		return models.Board{}, err
	}

	return bs.db.BoardRepository.GetByID(board.ID)
}

func (bs *BoardService) DeleteBoard(id uint) error {
	return bs.db.BoardRepository.Delete(id)
}

func (bs *BoardService) GetBoard(id uint) (models.Board, error) {
	board, err := bs.db.BoardRepository.GetByID(id)
	if err != nil {
		return models.Board{}, err
	}

	bs.SortBoard(&board)

	return board, nil
}

func (bs *BoardService) GetBoardBySlug(slug string) (models.Board, error) {
	board, err := bs.db.BoardRepository.GetFirst(repository.WithWhere("slug = ?", slug))
	if err != nil {
		return models.Board{}, err
	}

	bs.SortBoard(&board)

	return board, nil
}

func (bs *BoardService) GetBoardWithPermissions(userID, boardID uint) (models.Board, error) {
	can, err := bs.rs.CheckRole(userID, boardID, MemberRole)
	if err != nil {
		return models.Board{}, err
	}

	if !can {
		return models.Board{}, errors.New("forbidden")
	}

	board, err := bs.GetBoard(boardID)
	if err != nil {
		return models.Board{}, err
	}

	return board, nil
}

func (bs *BoardService) GetBoardBySlugWithPermissions(userID uint, slug string) (models.Board, error) {
	board, err := bs.GetBoardBySlug(slug)
	if err != nil {
		return models.Board{}, err
	}

	can, err := bs.rs.CheckRole(userID, board.ID, MemberRole)
	if err != nil {
		return models.Board{}, err
	}

	if !can {
		return models.Board{}, errors.New("forbidden")
	}

	return board, nil
}

func (bs *BoardService) SortBoard(board *models.Board) {
	sort.Slice(board.Swimlanes, func(i, j int) bool { return board.Swimlanes[i].Order < board.Swimlanes[j].Order })
	sort.Slice(board.Columns, func(i, j int) bool { return board.Columns[i].Order < board.Columns[j].Order })
}

func (bs *BoardService) ListBoards(userID uint) ([]models.Board, error) {
	MemberRole, err := bs.db.BoardRoleRepository.GetFirst(repository.WithWhere("name = ?", MemberRole))
	if err != nil {
		return []models.Board{}, err
	}

	userBoardRoles, err := bs.db.UserBoardRoleRepository.GetAll(repository.WithWhere("user_id = ? AND id = ?", userID, MemberRole.ID), repository.WithPreload("Board"))
	if err != nil {
		return []models.Board{}, err
	}

	boards := make([]models.Board, len(userBoardRoles))
	for i, userBoardRole := range userBoardRoles {
		boards[i] = userBoardRole.Board
	}

	return boards, nil
}
