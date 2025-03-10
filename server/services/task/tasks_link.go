package task

import (
	"errors"
	"server/database/repository"
	"server/models"
	"server/services/role"
)

func (ts *TaskService) GetTaskLink(linkID uint) (models.TaskLinks, error) {
	link, err := ts.db.TaskLinkRepository.GetByID(linkID, repository.WithPreload("SrcTask"), repository.WithPreload("DstTask"), repository.WithPreload("SrcTask.Board"), repository.WithPreload("DstTask.Board"))
	if err != nil {
		return models.TaskLinks{}, err
	}
	return link, nil
}

func (ts *TaskService) CreateTaskLink(userID uint, srcTaskID uint, dstTaskID uint, linkType string) (models.TaskLinks, error) {
	srcTask, err := ts.db.TaskRepository.GetByID(srcTaskID)
	if err != nil {
		return models.TaskLinks{}, err
	}
	if can, _ := ts.rs.CheckRole(userID, srcTask.BoardID, role.MemberRole); !can {
		return models.TaskLinks{}, errors.New("forbidden")
	}

	dstTask, err := ts.db.TaskRepository.GetByID(dstTaskID)
	if err != nil {
		return models.TaskLinks{}, err
	}
	if can, _ := ts.rs.CheckRole(userID, dstTask.BoardID, role.MemberRole); !can {
		return models.TaskLinks{}, errors.New("forbidden")
	}

	lType, ok := repository.LinkTypeMap[linkType]
	inverseLType, inverseOk := repository.InverseLinkTypeMap[linkType]
	if !ok && !inverseOk {
		return models.TaskLinks{}, errors.New("invalid link type")
	}

	rSrcId := srcTaskID
	rDstId := dstTaskID
	rLType := lType
	if !ok {
		rSrcId = dstTaskID
		rDstId = srcTaskID
		rLType = inverseLType
	}

	link := models.TaskLinks{
		SrcTaskID: rSrcId,
		DstTaskID: rDstId,
		LinkType:  string(rLType),
	}
	if err = ts.db.TaskLinkRepository.Create(&link); err != nil {
		return models.TaskLinks{}, err
	}
	return ts.GetTaskLink(link.ID)
}

func (ts *TaskService) DeleteTaskLinkRequest(userID uint, linkID uint) (models.TaskLinks, error) {
	link, err := ts.db.TaskLinkRepository.GetByID(linkID, repository.WithPreload("SrcTask"), repository.WithPreload("DstTask"), repository.WithPreload("SrcTask.Board"), repository.WithPreload("DstTask.Board"))
	if err != nil {
		return models.TaskLinks{}, err
	}

	if can, _ := ts.rs.CheckRole(userID, link.SrcTask.BoardID, role.MemberRole); !can {
		return models.TaskLinks{}, errors.New("forbidden")
	}
	if can, _ := ts.rs.CheckRole(userID, link.DstTask.BoardID, role.MemberRole); !can {
		return models.TaskLinks{}, errors.New("forbidden")
	}

	err = ts.DeleteTaskLink(link.ID)
	if err != nil {
		return models.TaskLinks{}, err
	}
	return link, nil
}

func (ts *TaskService) DeleteTaskLink(linkID uint) error {
	return ts.db.TaskLinkRepository.Delete(linkID)
}
