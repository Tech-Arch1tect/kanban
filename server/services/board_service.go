package services

import (
	"errors"
	"server/config"
	"server/database/repository"
	"server/models"
	"sort"

	"server/internal/email"

	"gorm.io/gorm"
)

type BoardService struct {
	db  *repository.Database
	rs  *RoleService
	es  *email.EmailService
	cfg *config.Config
}

func NewBoardService(db *repository.Database, rs *RoleService, es *email.EmailService, cfg *config.Config) *BoardService {
	return &BoardService{
		db:  db,
		rs:  rs,
		es:  es,
		cfg: cfg,
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
	user, err := bs.db.UserRepository.GetByID(userID)
	if err != nil {
		return []models.Board{}, err
	}

	if user.Role == models.RoleAdmin {
		return bs.db.BoardRepository.GetAll(repository.WithPreload("Swimlanes", "Columns"))
	}

	MemberRole, err := bs.db.BoardRoleRepository.GetFirst(repository.WithWhere("name = ?", MemberRole))
	if err != nil {
		return []models.Board{}, err
	}

	userBoardRoles, err := bs.db.UserBoardRoleRepository.GetAll(repository.WithWhere("user_id = ? AND id = ?", userID, MemberRole.ID), repository.WithPreload("Board"), repository.WithPreload("Board.Swimlanes"), repository.WithPreload("Board.Columns"))
	if err != nil {
		return []models.Board{}, err
	}

	boards := make([]models.Board, len(userBoardRoles))
	for i, userBoardRole := range userBoardRoles {
		boards[i] = userBoardRole.Board
	}

	return boards, nil
}

func (bs *BoardService) GetUsersWithAccess(userID, boardID uint) ([]models.User, error) {
	can, err := bs.rs.CheckRole(userID, boardID, MemberRole)
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

func (bs *BoardService) AddOrInviteUserToBoard(authUserID, boardID uint, email string, role AppRole) error {
	authUser, err := bs.db.UserRepository.GetFirst(repository.WithWhere("id = ?", authUserID))
	if err != nil {
		return err
	}

	can, err := bs.rs.CheckRole(authUser.ID, boardID, AdminRole)
	if err != nil {
		return err
	}

	if !can {
		return errors.New("forbidden")
	}

	user, err := bs.db.UserRepository.GetFirst(repository.WithWhere("email = ?", email))
	if err == gorm.ErrRecordNotFound {
		return bs.InviteUserToBoard(boardID, email, role.Name)
	}

	return bs.AddUserToBoard(boardID, user.ID, role)
}

func (bs *BoardService) AddUserToBoard(boardID, userID uint, role AppRole) error {
	return bs.rs.AssignRole(userID, boardID, role)
}

func (bs *BoardService) InviteUserToBoard(boardID uint, email string, roleName string) error {
	board, err := bs.db.BoardRepository.GetByID(boardID)
	if err != nil {
		return err
	}

	boardInvite := models.BoardInvite{
		BoardID:  boardID,
		Email:    email,
		RoleName: roleName,
	}

	err = bs.db.BoardInviteRepository.Create(&boardInvite)
	if err != nil {
		return err
	}

	err = bs.es.SendHTMLTemplate(email, "Invite to Board", "inviteToBoard.tmpl", map[string]string{
		"boardName": board.Name,
		"appUrl":    bs.cfg.AppUrl,
		"appName":   bs.cfg.AppName,
	})
	if err != nil {
		return err
	}

	return nil
}
