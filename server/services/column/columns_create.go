package column

import (
	"errors"
	"server/database/repository"
	"server/models"
	"server/services/role"
)

func (cs *ColumnService) CreateColumn(userID uint, boardID uint, name string) (models.Column, error) {
	can, err := cs.rs.CheckRole(userID, boardID, role.AdminRole)
	if err != nil {
		return models.Column{}, err
	}

	if !can {
		return models.Column{}, errors.New("forbidden")
	}

	highestOrder, err := cs.db.ColumnRepository.GetFirst(
		repository.WithWhere("board_id = ?", boardID),
		repository.WithOrder("`order` DESC"),
	)
	if err != nil {
		return models.Column{}, err
	}

	nextOrder := highestOrder.Order + 1

	column := models.Column{
		BoardID: boardID,
		Name:    name,
		Order:   nextOrder,
	}

	err = cs.db.ColumnRepository.Create(&column)
	if err != nil {
		return models.Column{}, err
	}

	return column, nil
}
