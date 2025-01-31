package column

import (
	"errors"
	"server/database/repository"
	"server/models"
	"server/services/role"
)

type ColumnService struct {
	db *repository.Database
	rs *role.RoleService
}

func NewColumnService(db *repository.Database, rs *role.RoleService) *ColumnService {
	return &ColumnService{db: db, rs: rs}
}

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

func (cs *ColumnService) DeleteColumn(userID uint, columnID uint) (models.Column, error) {
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

	err = cs.db.ColumnRepository.Delete(column.ID)
	if err != nil {
		return models.Column{}, err
	}

	return column, nil
}

func (cs *ColumnService) EditColumn(userID uint, columnID uint, name string) (models.Column, error) {
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

	columnMap := make(map[uint]int)
	for i, c := range allBoardColumns {
		columnMap[c.ID] = i
	}

	currentIdx, relativeIdx := columnMap[column.ID], columnMap[relativeColumn.ID]
	targetIdx := relativeIdx
	if direction == "before" {
		targetIdx++
	}

	allBoardColumns = append(allBoardColumns[:currentIdx], allBoardColumns[currentIdx+1:]...)

	if currentIdx < targetIdx {
		targetIdx--
	}

	allBoardColumns = append(allBoardColumns[:targetIdx], append([]models.Column{column}, allBoardColumns[targetIdx:]...)...)

	for i, col := range allBoardColumns {
		col.Order = i + 1
		if err := cs.db.ColumnRepository.Update(&col); err != nil {
			return models.Column{}, errors.New("failed to update column order")
		}
	}

	return column, nil
}
