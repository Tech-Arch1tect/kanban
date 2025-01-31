package swimlane

import (
	"errors"
	"server/database/repository"
	"server/models"
	"server/services/role"
)

type SwimlaneService struct {
	db *repository.Database
	rs *role.RoleService
}

func NewSwimlaneService(db *repository.Database, rs *role.RoleService) *SwimlaneService {
	return &SwimlaneService{db: db, rs: rs}
}

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

func (ss *SwimlaneService) MoveSwimlane(userID uint, swimlaneID uint, relativeID uint, direction string) (models.Swimlane, error) {
	swimlane, err := ss.db.SwimlaneRepository.GetByID(swimlaneID)
	if err != nil {
		return models.Swimlane{}, err
	}

	relativeSwimlane, err := ss.db.SwimlaneRepository.GetByID(relativeID)
	if err != nil {
		return models.Swimlane{}, err
	}

	if swimlane.BoardID != relativeSwimlane.BoardID {
		return models.Swimlane{}, errors.New("swimlanes are not on the same board")
	}

	can, err := ss.rs.CheckRole(userID, swimlane.BoardID, role.AdminRole)
	if err != nil {
		return models.Swimlane{}, err
	}

	if !can {
		return models.Swimlane{}, errors.New("forbidden")
	}

	allBoardSwimlanes, err := ss.db.SwimlaneRepository.GetAll(repository.WithWhere("board_id = ?", swimlane.BoardID))
	if err != nil {
		return models.Swimlane{}, err
	}

	swimlaneMap := make(map[uint]int)
	for i, c := range allBoardSwimlanes {
		swimlaneMap[c.ID] = i
	}

	currentIdx, relativeIdx := swimlaneMap[swimlane.ID], swimlaneMap[relativeSwimlane.ID]
	targetIdx := relativeIdx
	if direction == "before" {
		targetIdx++
	}

	allBoardSwimlanes = append(allBoardSwimlanes[:currentIdx], allBoardSwimlanes[currentIdx+1:]...)

	if currentIdx < targetIdx {
		targetIdx--
	}

	allBoardSwimlanes = append(allBoardSwimlanes[:targetIdx], append([]models.Swimlane{swimlane}, allBoardSwimlanes[targetIdx:]...)...)

	for i, swimlane := range allBoardSwimlanes {
		swimlane.Order = i + 1
		if err := ss.db.SwimlaneRepository.Update(&swimlane); err != nil {
			return models.Swimlane{}, errors.New("failed to update swimlane order")
		}
	}

	return swimlane, nil
}
