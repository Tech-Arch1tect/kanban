package notification

import (
	"net/http"
	"server/internal/helpers"
	"server/models"
	"server/services/notification"

	"github.com/gin-gonic/gin"
)

type NotificationController struct {
	notificationService *notification.NotificationService
	hs                  *helpers.HelperService
}

func NewNotificationController(notificationService *notification.NotificationService, hs *helpers.HelperService) *NotificationController {
	return &NotificationController{notificationService: notificationService, hs: hs}
}

type CreateNotificationRequest struct {
	Name         string   `json:"name" binding:"required"`
	Method       string   `json:"method" binding:"required,oneof=webhook email"`
	WebhookURL   string   `json:"webhook_url"`
	Email        string   `json:"email"`
	Events       []string `json:"events" binding:"required"`
	Boards       []uint   `json:"boards" binding:"required"`
	OnlyAssignee bool     `json:"only_assignee"`
}

type CreateNotificationResponse struct {
	ID uint `json:"id"`
}

// @Summary Create a notification configuration
// @Description Create a notification configuration
// @Tags notifications
// @Security cookieAuth
// @Security csrf
// @Accept json
// @Produce json
// @Param request body CreateNotificationRequest true "Notification details"
// @Success 200 {object} CreateNotificationResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Failure 403 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/notifications/create [post]
func (nc *NotificationController) CreateNotification(c *gin.Context) {
	var request CreateNotificationRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := nc.hs.GetUserFromSession(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	if len(request.Boards) == 0 || len(request.Events) == 0 || (request.WebhookURL == "" && request.Email == "") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Boards, events and (webhook_url or email) are required"})
		return
	}

	notification, err := nc.notificationService.CreateNotification(&user, request.Name, request.Method, request.WebhookURL, request.Email, request.Events, request.Boards, request.OnlyAssignee)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": notification.ID})
}

type GetEventsResponse struct {
	Events []string `json:"events"`
}

// @Summary Get available notification events
// @Description Retrieve the list of available event types for notifications.
// @Tags notifications
// @Security cookieAuth
// @Produce json
// @Success 200 {object} GetEventsResponse
// @Failure 401 {object} models.ErrorResponse
// @Router /api/v1/notifications/events [get]
func (nc *NotificationController) GetEvents(c *gin.Context) {
	events := nc.notificationService.GetEvents()
	c.JSON(http.StatusOK, GetEventsResponse{Events: events})
}

type GetNotificationsResponse struct {
	Notifications []models.NotificationConfiguration `json:"notifications"`
}

// @Summary Get all notification configurations
// @Description Retrieve the list of all notification configurations.
// @Tags notifications
// @Security cookieAuth
// @Produce json
// @Success 200 {object} GetNotificationsResponse
// @Failure 401 {object} models.ErrorResponse
// @Router /api/v1/notifications/list [get]
func (nc *NotificationController) GetNotifications(c *gin.Context) {
	user, err := nc.hs.GetUserFromSession(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	notifications, err := nc.notificationService.GetNotificationConfigurations(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, GetNotificationsResponse{Notifications: notifications})
}

type DeleteNotificationRequest struct {
	ID uint `json:"id" binding:"required"`
}

type DeleteNotificationResponse struct {
	ID uint `json:"id"`
}

// @Summary Delete a notification configuration
// @Description Delete a notification configuration
// @Tags notifications
// @Security cookieAuth
// @Security csrf
// @Accept json
// @Produce json
// @Param request body DeleteNotificationRequest true "Notification ID"
// @Success 200 {object} DeleteNotificationResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Failure 403 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/notifications/delete [post]
func (nc *NotificationController) DeleteNotification(c *gin.Context) {
	var request DeleteNotificationRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := nc.hs.GetUserFromSession(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	err = nc.notificationService.DeleteNotification(&user, request.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, DeleteNotificationResponse(request))
}
