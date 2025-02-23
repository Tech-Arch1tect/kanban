package swimlane

import (
	"net/http"
	"server/database/repository"
	"server/internal/helpers"
	"server/services/role"
	"server/services/swimlane"

	"github.com/gin-gonic/gin"
)

type SwimlaneController struct {
	db *repository.Database
	cs *swimlane.SwimlaneService
	rs *role.RoleService
	hs *helpers.HelperService
}

func NewSwimlaneController(db *repository.Database, cs *swimlane.SwimlaneService, rs *role.RoleService, hs *helpers.HelperService) *SwimlaneController {
	return &SwimlaneController{db: db, cs: cs, rs: rs, hs: hs}
}

// @Summary Create a swimlane
// @Description Create a swimlane for a board
// @Tags swimlanes
// @Security cookieAuth
// @Security csrf
// @Accept json
// @Produce json
// @Param request body CreateSwimlaneRequest true "Swimlane details"
// @Success 200 {object} CreateSwimlaneResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Failure 403 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/swimlanes/create [post]
func (cc *SwimlaneController) CreateSwimlane(c *gin.Context) {
	var request CreateSwimlaneRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := cc.hs.GetUserFromSession(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	swimlane, err := cc.cs.CreateSwimlane(user.ID, request.BoardID, request.Name)
	if err != nil {
		if err.Error() == "forbidden" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, CreateSwimlaneResponse{Swimlane: swimlane})
}

// @Summary Delete a swimlane
// @Description Delete a swimlane by ID
// @Tags swimlanes
// @Security cookieAuth
// @Security csrf
// @Accept json
// @Produce json
// @Param request body DeleteSwimlaneRequest true "Swimlane ID"
// @Success 200 {object} DeleteSwimlaneResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Failure 403 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/swimlanes/delete [post]
func (cc *SwimlaneController) DeleteSwimlane(c *gin.Context) {
	var request DeleteSwimlaneRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := cc.hs.GetUserFromSession(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	swimlane, err := cc.cs.DeleteSwimlaneRequest(user.ID, request.ID)
	if err != nil {
		if err.Error() == "forbidden" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, DeleteSwimlaneResponse{Swimlane: swimlane})
}

// @Summary Rename a swimlane
// @Description Rename a swimlane by ID
// @Tags swimlanes
// @Security cookieAuth
// @Security csrf
// @Accept json
// @Produce json
// @Param request body RenameSwimlaneRequest true "Swimlane details"
// @Success 200 {object} RenameSwimlaneResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Failure 403 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/swimlanes/rename [post]
func (cc *SwimlaneController) RenameSwimlane(c *gin.Context) {
	var request RenameSwimlaneRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := cc.hs.GetUserFromSession(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	swimlane, err := cc.cs.RenameSwimlane(user.ID, request.ID, request.Name)
	if err != nil {
		if err.Error() == "forbidden" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, RenameSwimlaneResponse{Swimlane: swimlane})
}

// @Summary Move a swimlane
// @Description Move a swimlane relative to another swimlane
// @Tags swimlanes
// @Security cookieAuth
// @Security csrf
// @Accept json
// @Produce json
// @Param request body MoveSwimlaneRequest true "Swimlane ID and relative swimlane ID and direction"
// @Success 200 {object} MoveSwimlaneResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Failure 403 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/swimlanes/move [post]
func (cc *SwimlaneController) MoveSwimlane(c *gin.Context) {
	var request MoveSwimlaneRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	user, err := cc.hs.GetUserFromSession(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	swimlane, err := cc.cs.MoveSwimlane(user.ID, request.ID, request.RelativeID, request.Direction)
	if err != nil {
		if err.Error() == "forbidden" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
		} else if err.Error() == "swimlanes are not on the same board" {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, MoveSwimlaneResponse{Swimlane: swimlane})
}
