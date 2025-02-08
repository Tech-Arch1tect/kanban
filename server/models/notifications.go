package models

import "time"

type NotificationConfiguration struct {
	Model
	Name         string              `gorm:"not null" json:"name"`
	UserID       uint                `gorm:"not null" json:"user_id"`
	User         User                `gorm:"foreignKey:UserID" json:"user"`
	Method       string              `gorm:"not null" json:"method"` // e.g. "email", "webhook"
	WebhookURL   string              `json:"webhook_url"`            // required if Method=="webhook"
	Email        string              `json:"email"`                  // required if Method=="email"
	Events       []NotificationEvent `gorm:"many2many:notification_configuration_events;" json:"events"`
	Boards       []Board             `gorm:"many2many:notification_configuration_boards;" json:"boards"`
	OnlyAssignee bool                `json:"only_assignee"`
}

type NotificationEvent struct {
	Model
	Name                       string                      `gorm:"unique;not null" json:"name"`
	NotificationConfigurations []NotificationConfiguration `gorm:"many2many:notification_configuration_events;" json:"notification_configurations"`
}

type NotificationLog struct {
	Model
	ConfigID    uint                      `gorm:"not null" json:"config_id"`
	Config      NotificationConfiguration `gorm:"foreignKey:ConfigID" json:"config"`
	EventName   string                    `gorm:"not null" json:"event_name"`
	Title       string                    `gorm:"not null" json:"title"`
	Description string                    `gorm:"not null" json:"description"`
	SentAt      time.Time                 `json:"sent_at"`
}
