package swimlane

import (
	"errors"
	"server/database/repository"
	"server/models"
	"server/services/role"
)

func (ss *SwimlaneService) CreateSwimlane(userID uint, boardID uint, name string) (models.Swimlane, error) {
	can, err := ss.rs.CheckRole(userID, boardID, role.AdminRole)
	if err != nil {
		return models.Swimlane{}, err
	}

	if !can {
		return models.Swimlane{}, errors.New("forbidden")
	}

	highestOrder, err := ss.db.SwimlaneRepository.GetFirst(
		repository.WithWhere("board_id = ?", boardID),
		repository.WithOrder("`order` DESC"),
	)
	if err != nil {
		return models.Swimlane{}, err
	}

	nextOrder := highestOrder.Order + 1

	swimlane := models.Swimlane{
		BoardID: boardID,
		Name:    name,
		Order:   nextOrder,
	}

	err = ss.db.SwimlaneRepository.Create(&swimlane)
	if err != nil {
		return models.Swimlane{}, err
	}

	return swimlane, nil
}
