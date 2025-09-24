package handler

import (
	"NWUCA-Management-System/server/internal/dto"
	"NWUCA-Management-System/server/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DepartmentHandler struct {
	service service.DepartmentService
}

func NewDepartmentHandler(service service.DepartmentService) *DepartmentHandler {
	return &DepartmentHandler{service: service}
}

// CreateDepartment
// @Summary Create a new department
// @Description Create a new department
// @Tags departments
// @Accept  json
// @Produce  json
// @Param   department     body    dto.CreateDepartmentRequest     true        "Department creation info"
// @Success 201 {object} model.Department
// @Failure 400 {object} dto.ErrorResponse "Invalid request body"
// @Failure 500 {object} dto.ErrorResponse "Failed to create department"
// @Security ApiKeyAuth
// @Router /departments [post]
func (h *DepartmentHandler) CreateDepartment(c *gin.Context) {
	var req dto.CreateDepartmentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	department, err := h.service.Create(req.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create department"})
		return
	}
	c.JSON(http.StatusCreated, department)
}

// GetDepartments
// @Summary Get all departments
// @Description Get a list of all departments
// @Tags departments
// @Produce  json
// @Success 200 {array} model.Department
// @Failure 500 {object} dto.ErrorResponse "Failed to get departments"
// @Security ApiKeyAuth
// @Router /departments [get]
func (h *DepartmentHandler) GetDepartments(c *gin.Context) {
	departments, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get departments"})
		return
	}
	c.JSON(http.StatusOK, departments)
}

// UpdateDepartment
// @Summary Update an existing department
// @Description Update an existing department by ID
// @Tags departments
// @Accept  json
// @Produce  json
// @Param   id     path    int     true        "Department ID"
// @Param   department     body    dto.UpdateDepartmentRequest     true        "Department update info"
// @Success 200 {object} model.Department
// @Failure 400 {object} dto.ErrorResponse "Invalid request body or ID"
// @Failure 500 {object} dto.ErrorResponse "Failed to update department"
// @Security ApiKeyAuth
// @Router /departments/{id} [put]
func (h *DepartmentHandler) UpdateDepartment(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid department ID"})
		return
	}
	var req dto.UpdateDepartmentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	department, err := h.service.Update(uint(id), req.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update department"})
		return
	}
	c.JSON(http.StatusOK, department)
}

// DeleteDepartment
// @Summary Delete a department
// @Description Delete a department by ID
// @Tags departments
// @Produce  json
// @Param   id     path    int     true        "Department ID"
// @Success 204 "No Content"
// @Failure 400 {object} dto.ErrorResponse "Invalid department ID"
// @Failure 500 {object} dto.ErrorResponse "Failed to delete department"
// @Security ApiKeyAuth
// @Router /departments/{id} [delete]
func (h *DepartmentHandler) DeleteDepartment(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid department ID"})
		return
	}
	if err := h.service.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete department"})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
