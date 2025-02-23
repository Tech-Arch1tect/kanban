package swimlane

import "server/models"

type CreateSwimlaneRequest struct {
	BoardID uint   `json:"board_id"`
	Name    string `json:"name"`
}

type CreateSwimlaneResponse struct {
	Swimlane models.Swimlane `json:"swimlane"`
}

type DeleteSwimlaneRequest struct {
	ID uint `json:"id"`
}

type DeleteSwimlaneResponse struct {
	Swimlane models.Swimlane `json:"swimlane"`
}

type RenameSwimlaneRequest struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type RenameSwimlaneResponse struct {
	Swimlane models.Swimlane `json:"swimlane"`
}

type MoveSwimlaneRequest struct {
	ID         uint   `json:"id" binding:"required"`
	RelativeID uint   `json:"relative_id" binding:"required"`
	Direction  string `json:"direction" binding:"required,oneof=before after"`
}

type MoveSwimlaneResponse struct {
	Swimlane models.Swimlane `json:"swimlane"`
}
