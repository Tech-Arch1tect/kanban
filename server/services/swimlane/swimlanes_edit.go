package swimlane

import (
	"errors"
	"server/models"
	"server/services/role"
)

func (ss *SwimlaneService) EditSwimlane(userID uint, swimlaneID uint, name string) (models.Swimlane, error) {
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

	swimlane.Name = name

	err = ss.db.SwimlaneRepository.Update(&swimlane)
	if err != nil {
		return models.Swimlane{}, err
	}

	return swimlane, nil
}
