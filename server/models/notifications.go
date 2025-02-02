package models

import "time"

type NotificationConfiguration struct {
	Model
	UserID     uint                `gorm:"not null"`
	User       User                `gorm:"foreignKey:UserID"`
	Method     string              `gorm:"not null"` // e.g. "email", "webhook"
	WebhookURL string              // required if Method=="webhook"
	Email      string              // required if Method=="email"
	Events     []NotificationEvent `gorm:"many2many:notification_configuration_events;"`
}

type NotificationEvent struct {
	Model
	Name                       string                      `gorm:"unique;not null"`
	NotificationConfigurations []NotificationConfiguration `gorm:"many2many:notification_configuration_events;"`
}

type NotificationLog struct {
	Model
	ConfigID    uint                      `gorm:"not null"`
	Config      NotificationConfiguration `gorm:"foreignKey:ConfigID"`
	EventName   string                    `gorm:"not null"`
	Title       string                    `gorm:"not null"`
	Description string                    `gorm:"not null"`
	SentAt      time.Time
}
