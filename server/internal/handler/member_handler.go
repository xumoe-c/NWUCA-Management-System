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

type MemberHandler struct {
	service service.MemberService
}

func NewMemberHandler(service service.MemberService) *MemberHandler {
	return &MemberHandler{service: service}
}

// CreateMember
// @Summary Create a new member
// @Description Create a new member
// @Tags members
// @Accept  json
// @Produce  json
// @Param   member     body    dto.CreateMemberRequest     true        "Member creation info"
// @Success 201 {object} model.Member
// @Failure 400 {object} dto.ErrorResponse "Invalid request body"
// @Failure 500 {object} dto.ErrorResponse "Failed to create member"
// @Security ApiKeyAuth
// @Router /members [post]
func (h *MemberHandler) CreateMember(c *gin.Context) {
	var req dto.CreateMemberRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{
			Code: http.StatusBadRequest,
			Msg:  "参数错误",
			Data: nil,
		})
		return
	}
	// The service layer might need to be adapted to take the DTO
	member, err := h.service.CreateMember(req)
	if err != nil {
		switch {
		case errors.Is(err, apperrors.ErrNotFound):
			c.JSON(http.StatusNotFound, dto.Response{
				Code: http.StatusNotFound,
				Msg:  "用户邮箱不存在",
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
		Data: member,
	})
}

// GetMembers
// @Summary Get all members
// @Description Get a list of all members
// @Tags members
// @Produce  json
// @Success 200 {array} model.Member
// @Failure 500 {object} dto.ErrorResponse "Failed to get members"
// @Security ApiKeyAuth
// @Router /members [get]
func (h *MemberHandler) GetMembers(c *gin.Context) {
	members, err := h.service.GetAllMembers()
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
		Data: members,
	})
}

// UpdateMember
// @Summary Update an existing member
// @Description Update an existing member by ID
// @Tags members
// @Accept  json
// @Produce  json
// @Param   id     path    int     true        "Member ID"
// @Param   member     body    dto.UpdateMemberRequest     true        "Member update info"
// @Success 200 {object} model.Member
// @Failure 400 {object} dto.ErrorResponse "Invalid request body or ID"
// @Failure 500 {object} dto.ErrorResponse "Failed to update member"
// @Security ApiKeyAuth
// @Router /members/{id} [put]
func (h *MemberHandler) UpdateMember(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{
			Code: http.StatusBadRequest,
			Msg:  "参数错误: 非法的id",
			Data: nil,
		})
		return
	}
	var req dto.UpdateMemberRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{
			Code: http.StatusBadRequest,
			Msg:  "参数错误",
			Data: nil,
		})
		return
	}
	// The service layer might need to be adapted to take the DTO
	member, err := h.service.UpdateMember(uint(id), req)
	if err != nil {
		switch {
		case errors.Is(err, apperrors.ErrNotFound):
			c.JSON(http.StatusNotFound, dto.Response{
				Code: http.StatusNotFound,
				Msg:  "会员不存在",
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
		Data: member,
	})
}

// DeleteMember
// @Summary Delete a member
// @Description Delete a member by ID
// @Tags members
// @Produce  json
// @Param   id     path    int     true        "Member ID"
// @Success 204 "No Content"
// @Failure 400 {object} dto.ErrorResponse "Invalid member ID"
// @Failure 500 {object} dto.ErrorResponse "Failed to delete member"
// @Security ApiKeyAuth
// @Router /members/{id} [delete]
func (h *MemberHandler) DeleteMember(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{
			Code: http.StatusBadRequest,
			Msg:  "参数错误: 非法的id",
			Data: nil,
		})
		return
	}
	if err := h.service.DeleteMember(uint(id)); err != nil {
		switch {
		case errors.Is(err, apperrors.ErrNotFound):
			c.JSON(http.StatusNotFound, dto.Response{
				Code: http.StatusNotFound,
				Msg:  "会员不存在",
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
		Msg:  "删除成功",
		Data: nil,
	})
}
