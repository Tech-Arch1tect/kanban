package comment

import "server/models"

type CreateCommentRequest struct {
	TaskID uint   `json:"task_id"`
	Text   string `json:"text"`
}

type CreateCommentResponse struct {
	Comment models.Comment `json:"comment"`
}

type DeleteCommentRequest struct {
	ID uint `json:"id"`
}

type DeleteCommentResponse struct {
	Comment models.Comment `json:"comment"`
}

type EditCommentRequest struct {
	ID   uint   `json:"id"`
	Text string `json:"text"`
}

type EditCommentResponse struct {
	Comment models.Comment `json:"comment"`
}
