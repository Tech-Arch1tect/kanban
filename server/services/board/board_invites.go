package board

import (
	"errors"
	"server/database/repository"
	"server/models"
	"server/services/role"
)

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

func (bs *BoardService) DeleteBoardInvite(boardInviteID uint) error {
	return bs.db.BoardInviteRepository.HardDelete(boardInviteID)
}

func (bs *BoardService) GetPendingInvites(userID, boardID uint) ([]models.BoardInvite, error) {
	can, err := bs.rs.CheckRole(userID, boardID, role.AdminRole)
	if err != nil || !can {
		return nil, errors.New("forbidden")
	}

	return bs.db.BoardInviteRepository.GetAll(repository.WithWhere("board_id = ?", boardID))
}

func (bs *BoardService) RemovePendingInviteRequest(userID, inviteID uint) (models.BoardInvite, error) {
	invite, err := bs.db.BoardInviteRepository.GetByID(inviteID)
	if err != nil {
		return models.BoardInvite{}, err
	}

	can, err := bs.rs.CheckRole(userID, invite.BoardID, role.AdminRole)
	if err != nil || !can {
		return models.BoardInvite{}, errors.New("forbidden")
	}

	return invite, bs.DeleteBoardInvite(inviteID)
}
