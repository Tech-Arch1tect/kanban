package column

import (
	"errors"
	"server/database/repository"
	"server/models"
	"server/services/role"
)

func (cs *ColumnService) MoveColumn(userID uint, columnID uint, relativeID uint, direction string) (models.Column, error) {
	column, err := cs.db.ColumnRepository.GetByID(columnID)
	if err != nil {
		return models.Column{}, err
	}

	relativeColumn, err := cs.db.ColumnRepository.GetByID(relativeID)
	if err != nil {
		return models.Column{}, err
	}

	if column.BoardID != relativeColumn.BoardID {
		return models.Column{}, errors.New("columns are not on the same board")
	}

	can, err := cs.rs.CheckRole(userID, column.BoardID, role.AdminRole)
	if err != nil {
		return models.Column{}, err
	}

	if !can {
		return models.Column{}, errors.New("forbidden")
	}

	allBoardColumns, err := cs.db.ColumnRepository.GetAll(repository.WithWhere("board_id = ?", column.BoardID))
	if err != nil {
		return models.Column{}, err
	}

	// Build a map of column IDs to their indexes.
	columnMap := make(map[uint]int)
	for i, c := range allBoardColumns {
		columnMap[c.ID] = i
	}

	currentIdx, relativeIdx := columnMap[column.ID], columnMap[relativeColumn.ID]
	targetIdx := relativeIdx
	if direction == "before" {
		targetIdx++
	}

	// Remove the column from its current position.
	allBoardColumns = append(allBoardColumns[:currentIdx], allBoardColumns[currentIdx+1:]...)

	if currentIdx < targetIdx {
		targetIdx--
	}

	// Insert the column at the new index.
	allBoardColumns = append(allBoardColumns[:targetIdx], append([]models.Column{column}, allBoardColumns[targetIdx:]...)...)

	// Update order for all columns.
	for i, col := range allBoardColumns {
		col.Order = i + 1
		if err := cs.db.ColumnRepository.Update(&col); err != nil {
			return models.Column{}, errors.New("failed to update column order")
		}
	}

	return column, nil
}
