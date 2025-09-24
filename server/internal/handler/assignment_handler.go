package handler

import (
	"NWUCA-Management-System/server/internal/dto"
	"NWUCA-Management-System/server/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AssignmentHandler struct {
	service service.AssignmentService
}

func NewAssignmentHandler(service service.AssignmentService) *AssignmentHandler {
	return &AssignmentHandler{service: service}
}

// CreateAssignment
// @Summary Create a new assignment
// @Description Create a new assignment
// @Tags assignments
// @Accept  json
// @Produce  json
// @Param   assignment     body    dto.CreateAssignmentRequest     true        "Assignment creation info"
// @Success 201 {object} model.Assignment
// @Failure 400 {object} dto.ErrorResponse "Invalid request body"
// @Failure 500 {object} dto.ErrorResponse "Failed to create assignment"
// @Security ApiKeyAuth
// @Router /assignments [post]
func (h *AssignmentHandler) CreateAssignment(c *gin.Context) {
	var req dto.CreateAssignmentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// The service layer might need to be adapted to take the DTO
	assignment, err := h.service.CreateAssignment(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create assignment"})
		return
	}
	c.JSON(http.StatusCreated, assignment)
}

// GetAssignments
// @Summary Get all assignments
// @Description Get a list of all assignments
// @Tags assignments
// @Produce  json
// @Success 200 {array} model.Assignment
// @Failure 500 {object} dto.ErrorResponse "Failed to get assignments"
// @Security ApiKeyAuth
// @Router /assignments [get]
func (h *AssignmentHandler) GetAssignments(c *gin.Context) {
	assignments, err := h.service.GetAllAssignments()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get assignments"})
		return
	}
	c.JSON(http.StatusOK, assignments)
}

// UpdateAssignment
// @Summary Update an existing assignment
// @Description Update an existing assignment by ID
// @Tags assignments
// @Accept  json
// @Produce  json
// @Param   id     path    int     true        "Assignment ID"
// @Param   assignment     body    dto.UpdateAssignmentRequest     true        "Assignment update info"
// @Success 200 {object} model.Assignment
// @Failure 400 {object} dto.ErrorResponse "Invalid request body or ID"
// @Failure 500 {object} dto.ErrorResponse "Failed to update assignment"
// @Security ApiKeyAuth
// @Router /assignments/{id} [put]
func (h *AssignmentHandler) UpdateAssignment(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid assignment ID"})
		return
	}
	var req dto.UpdateAssignmentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// The service layer might need to be adapted to take the DTO
	assignment, err := h.service.UpdateAssignment(uint(id), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update assignment"})
		return
	}
	c.JSON(http.StatusOK, assignment)
}

// DeleteAssignment
// @Summary Delete an assignment
// @Description Delete an assignment by ID
// @Tags assignments
// @Produce  json
// @Param   id     path    int     true        "Assignment ID"
// @Success 204 "No Content"
// @Failure 400 {object} dto.ErrorResponse "Invalid assignment ID"
// @Failure 500 {object} dto.ErrorResponse "Failed to delete assignment"
// @Security ApiKeyAuth
// @Router /assignments/{id} [delete]
func (h *AssignmentHandler) DeleteAssignment(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid assignment ID"})
		return
	}
	if err := h.service.DeleteAssignment(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete assignment"})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
