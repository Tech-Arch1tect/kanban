package notification

import (
	"errors"
	"server/models"
)

func (ns *NotificationSubscriber) SendNotification(subject, body string, notif models.NotificationConfiguration) error {
	if notif.Method == "email" {
		err := ns.email.SendPlainText(notif.Email, subject, body)
		if err != nil {
			return err
		}
	} else if notif.Method == "webhook" {
		return errors.New("webhook method not implemented")
	}
	return nil
}
