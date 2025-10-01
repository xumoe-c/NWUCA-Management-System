package handler

import (
	"NWUCA-Management-System/server/internal/dto"
	apperrors "NWUCA-Management-System/server/internal/errors"
	"NWUCA-Management-System/server/internal/service"
	"errors"
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
		c.JSON(http.StatusBadRequest, dto.Response{
			Code: http.StatusBadRequest,
			Msg:  "参数错误",
			Data: nil,
		})
		return
	}
	department, err := h.service.Create(req.Name)
	if err != nil {
		switch {
		case errors.Is(err, apperrors.ErrDepartmentExists):
			c.JSON(http.StatusConflict, dto.Response{
				Code: http.StatusConflict,
				Msg:  "名称被占用",
				Data: nil,
			})
		default:
			c.JSON(http.StatusInternalServerError, dto.Response{
				Code: http.StatusInternalServerError,
				Msg:  "服务器内部错误",
				Data: nil,
			})
		}
		return
	}
	c.JSON(http.StatusCreated, dto.Response{
		Code: http.StatusCreated,
		Msg:  "创建成功",
		Data: department,
	})
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
		c.JSON(http.StatusInternalServerError, dto.Response{
			Code: http.StatusInternalServerError,
			Msg:  "服务器内部错误",
			Data: nil,
		})
		return
	}
	c.JSON(http.StatusOK, dto.Response{
		Code: http.StatusOK,
		Msg:  "获取成功",
		Data: departments,
	})
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
		c.JSON(http.StatusBadRequest, dto.Response{
			Code: http.StatusBadRequest,
			Msg:  "参数错误: 非法的id",
			Data: nil,
		})
		return
	}
	var req dto.UpdateDepartmentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{
			Code: http.StatusBadRequest,
			Msg:  "参数错误",
			Data: nil,
		})
		return
	}
	department, err := h.service.Update(uint(id), req.Name)
	if err != nil {
		switch {
		case errors.Is(err, apperrors.ErrNotFound):
			c.JSON(http.StatusNotFound, dto.Response{
				Code: http.StatusNotFound,
				Msg:  "没找到",
				Data: nil,
			})
		default:
			c.JSON(http.StatusInternalServerError, dto.Response{
				Code: http.StatusInternalServerError,
				Msg:  "服务器内部错误",
				Data: nil,
			})
		}
		return
	}
	c.JSON(http.StatusOK, dto.Response{
		Code: http.StatusOK,
		Msg:  "更新成功",
		Data: department,
	})
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
		c.JSON(http.StatusBadRequest, dto.Response{
			Code: http.StatusBadRequest,
			Msg:  "参数错误: 非法的id",
			Data: nil,
		})
		return
	}
	if err := h.service.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{
			Code: http.StatusInternalServerError,
			Msg:  "服务器内部错误",
			Data: nil,
		})
		return
	}
	c.JSON(http.StatusOK, dto.Response{
		Code: http.StatusOK,
		Msg:  "删除成功",
		Data: nil,
	})
}
