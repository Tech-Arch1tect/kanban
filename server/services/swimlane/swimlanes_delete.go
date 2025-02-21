package swimlane

import (
	"errors"
	"server/models"
	"server/services/role"
)

func (ss *SwimlaneService) DeleteSwimlaneRequest(userID uint, swimlaneID uint) (models.Swimlane, error) {
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

	err = ss.DeleteSwimlane(swimlane.ID)
	if err != nil {
		return models.Swimlane{}, err
	}

	return swimlane, nil
}

func (ss *SwimlaneService) DeleteSwimlane(swimlaneID uint) error {
	return ss.db.SwimlaneRepository.Delete(swimlaneID)
}
