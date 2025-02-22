package column

import (
	"errors"
	"server/models"
	"server/services/role"
)

func (cs *ColumnService) RenameColumn(userID uint, columnID uint, name string) (models.Column, error) {
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

	column.Name = name

	err = cs.db.ColumnRepository.Update(&column)
	if err != nil {
		return models.Column{}, err
	}

	return column, nil
}
