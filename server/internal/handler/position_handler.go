package handler

import (
	"NWUCA-Management-System/server/internal/dto"
	"NWUCA-Management-System/server/internal/service"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PositionHandler struct {
	service service.PositionService
}

func NewPositionHandler(service service.PositionService) *PositionHandler {
	return &PositionHandler{service: service}
}

// CreatePosition
// @Summary Create a new position
// @Description Create a new position
// @Tags positions
// @Accept  json
// @Produce  json
// @Param   position     body    dto.CreatePositionRequest     true        "Position creation info"
// @Success 201 {object} model.Position
// @Failure 400 {object} dto.ErrorResponse "Invalid request body"
// @Failure 500 {object} dto.ErrorResponse "Failed to create position"
// @Security ApiKeyAuth
// @Router /positions [post]
func (h *PositionHandler) CreatePosition(c *gin.Context) {
	var req dto.CreatePositionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{
			Code: http.StatusBadRequest,
			Msg:  "参数错误",
			Data: nil,
		})
		return
	}
	position, err := h.service.Create(req.Name)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrNameExists):
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
		Data: position,
	})
}

// GetPositions
// @Summary Get all positions
// @Description Get a list of all positions
// @Tags positions
// @Produce  json
// @Success 200 {array} model.Position
// @Failure 500 {object} dto.ErrorResponse "Failed to get positions"
// @Security ApiKeyAuth
// @Router /positions [get]
func (h *PositionHandler) GetPositions(c *gin.Context) {
	positions, err := h.service.GetAll()
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
		Data: positions,
	})
}

// UpdatePosition
// @Summary Update an existing position
// @Description Update an existing position by ID
// @Tags positions
// @Accept  json
// @Produce  json
// @Param   id     path    int     true        "Position ID"
// @Param   position     body    dto.UpdatePositionRequest     true        "Position update info"
// @Success 200 {object} model.Position
// @Failure 400 {object} dto.ErrorResponse "Invalid request body or ID"
// @Failure 500 {object} dto.ErrorResponse "Failed to update position"
// @Security ApiKeyAuth
// @Router /positions/{id} [put]
func (h *PositionHandler) UpdatePosition(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{
			Code: http.StatusBadRequest,
			Msg:  "参数错误: 非法的id",
			Data: nil,
		})
		return
	}
	var req dto.UpdatePositionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{
			Code: http.StatusBadRequest,
			Msg:  "参数错误",
			Data: nil,
		})
		return
	}
	position, err := h.service.Update(uint(id), req.Name)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrPositionNotExists):
			c.JSON(http.StatusNotFound, dto.Response{
				Code: http.StatusNotFound,
				Msg:  "未找到",
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
		Msg:  "修改成功",
		Data: position,
	})
}

// DeletePosition
// @Summary Delete a position
// @Description Delete a position by ID
// @Tags positions
// @Produce  json
// @Param   id     path    int     true        "Position ID"
// @Success 204 "No Content"
// @Failure 400 {object} dto.ErrorResponse "Invalid position ID"
// @Failure 500 {object} dto.ErrorResponse "Failed to delete position"
// @Security ApiKeyAuth
// @Router /positions/{id} [delete]
func (h *PositionHandler) DeletePosition(c *gin.Context) {
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
