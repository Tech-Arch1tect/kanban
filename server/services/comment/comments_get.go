package comment

import (
	"server/database/repository"
	"server/models"
)

func (cs *CommentService) GetComment(commentID uint) (models.Comment, error) {
	comment, err := cs.db.CommentRepository.GetByID(commentID, repository.WithPreload("Task"), repository.WithPreload("User"))
	if err != nil {
		return models.Comment{}, err
	}

	return comment, nil
}
