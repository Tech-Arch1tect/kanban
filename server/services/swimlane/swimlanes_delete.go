package swimlane

import (
	"errors"
	"server/models"
	"server/services/role"
)

func (ss *SwimlaneService) DeleteSwimlane(userID uint, swimlaneID uint) (models.Swimlane, error) {
	swimlane, err := ss.db.SwimlaneRepository.GetByID(swimlaneID)
	if err != nil {
		return models.Swimlane{}, err
	}

	can, err := ss.rs.CheckRole(userID, swimlane.BoardID, role.AdminRole)
	if err != nil {
		return models.Swimlane{}, err
	}

	if !can {
		return models.Swimlane{}, errors.New("forbidden")
	}

	err = ss.db.SwimlaneRepository.Delete(swimlane.ID)
	if err != nil {
		return models.Swimlane{}, err
	}

	return swimlane, nil
}
