package comment

import (
	"errors"
	"server/models"
	"server/services/role"
)

func (cs *CommentService) CreateComment(userID, taskID uint, text string) (models.Comment, error) {
	task, err := cs.db.TaskRepository.GetByID(taskID)
	if err != nil {
		return models.Comment{}, err
	}

	can, err := cs.rs.CheckRole(userID, task.BoardID, role.MemberRole)
	if err != nil {
		return models.Comment{}, err
	}

	if !can {
		return models.Comment{}, errors.New("forbidden")
	}

	comment := models.Comment{
		TaskID: taskID,
		Text:   text,
		UserID: userID,
	}

	if err := cs.db.CommentRepository.Create(&comment); err != nil {
		return models.Comment{}, err
	}

	return comment, nil
}

func (cs *CommentService) CreateCommentReaction(userID, commentID uint, reactionStr string) (models.Reaction, error) {
	comment, err := cs.GetComment(commentID)
	if err != nil {
		return models.Reaction{}, err
	}

	can, err := cs.rs.CheckRole(userID, comment.Task.BoardID, role.MemberRole)
	if err != nil {
		return models.Reaction{}, err
	}

	if !can {
		return models.Reaction{}, errors.New("forbidden")
	}

	reaction := models.Reaction{
		CommentID: commentID,
		UserID:    userID,
		Reaction:  reactionStr,
	}

	if err := cs.db.CommentReactionRepository.Create(&reaction); err != nil {
		return models.Reaction{}, err
	}

	return reaction, nil
}
