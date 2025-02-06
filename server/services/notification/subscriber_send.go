package notification

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"server/models"
)

func (ns *NotificationSubscriber) SendNotification(event, subject, body string, notif models.NotificationConfiguration) error {
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

type WebhookPayload struct {
	Event   string `json:"event"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

func (ns *NotificationSubscriber) SendWebhookNotification(event, subject, body string, notif models.NotificationConfiguration) error {
	payload := WebhookPayload{
		Event:   event,
		Subject: subject,
		Body:    body,
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", notif.WebhookURL, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= http.StatusBadRequest {
		return errors.New("failed webhook call, status: " + resp.Status)
	}

	return nil
}
