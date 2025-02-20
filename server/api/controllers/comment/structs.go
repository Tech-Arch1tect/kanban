package comment

import "server/models"

type CreateCommentRequest struct {
	TaskID uint   `json:"task_id"`
	Text   string `json:"text"`
}

type CreateCommentResponse struct {
	Comment models.Comment `json:"comment"`
}

type CreateCommentReactionRequest struct {
	CommentID uint   `json:"comment_id"`
	Reaction  string `json:"reaction"`
}

type CreateCommentReactionResponse struct {
	Reaction models.Reaction `json:"reaction"`
}

type DeleteCommentReactionRequest struct {
	ReactionID uint `json:"reaction_id"`
}

type DeleteCommentReactionResponse struct {
	Reaction models.Reaction `json:"reaction"`
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
