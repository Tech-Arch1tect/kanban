package models

type Swimlane struct {
	Model
	BoardID uint   `json:"board_id"`
	Name    string `json:"name"`
	Order   int    `json:"order"`
}
