package services

import (
	"errors"
	"server/database/repository"
	"server/models"
	"sort"
)

type BoardService struct {
	db *repository.Database
	ps *PermissionService
}

func NewBoardService(db *repository.Database, ps *PermissionService) *BoardService {
	return &BoardService{
		db: db,
		ps: ps,
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
	can, err := bs.ps.CheckPermission(userID, boardID, ViewPermission)
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

	can, err := bs.ps.CheckPermission(userID, board.ID, ViewPermission)
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
	ViewPermission, err := bs.db.BoardPermissionRepository.GetFirst(repository.WithWhere("name = ?", ViewPermission))
	if err != nil {
		return []models.Board{}, err
	}

	userBoardPermissions, err := bs.db.UserBoardPermissionRepository.GetAll(repository.WithWhere("user_id = ? AND id = ?", userID, ViewPermission.ID), repository.WithPreload("Board"))
	if err != nil {
		return []models.Board{}, err
	}

	boards := make([]models.Board, len(userBoardPermissions))
	for i, userBoardPermission := range userBoardPermissions {
		boards[i] = userBoardPermission.Board
	}

	return boards, nil
}
