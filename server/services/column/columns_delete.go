package column

import (
	"errors"
	"server/models"
	"server/services/role"
)

func (cs *ColumnService) DeleteColumnRequest(userID uint, columnID uint) (models.Column, error) {
	column, err := cs.db.ColumnRepository.GetByID(columnID)
	if err != nil {
		return models.Column{}, err
	}

	can, err := cs.rs.CheckRole(userID, column.BoardID, role.AdminRole)
	if err != nil {
		return models.Column{}, err
	}

	if !can {
		return models.Column{}, errors.New("forbidden")
	}

	err = cs.DeleteColumn(column.ID)
	if err != nil {
		return models.Column{}, err
	}

	return column, nil
}

func (cs *ColumnService) DeleteColumn(columnID uint) error {
	return cs.db.ColumnRepository.Delete(columnID)
}
