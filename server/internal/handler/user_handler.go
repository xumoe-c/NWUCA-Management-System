package handler

import (
	"NWUCA-Management-System/server/internal/dto"
	"NWUCA-Management-System/server/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserHandler 包含用户相关的处理器
type UserHandler struct {
	userService service.UserService // 依赖服务层的接口
}

// NewUserHandler 创建一个 UserHandler 实例
func NewUserHandler(svc service.UserService) *UserHandler {
	return &UserHandler{userService: svc}
}

// Register
// @Summary Register a new user
// @Description Register a new user with username, email, and password
// @Tags users
// @Accept  json
// @Produce  json
// @Param   user     body    dto.RegisterRequest     true        "User registration info"
// @Success 201 {object} dto.UserResponse "User registered successfully"
// @Failure 400 {object} dto.ErrorResponse "Invalid request body"
// @Failure 500 {object} dto.ErrorResponse "Failed to register user"
// @Router /register [post]
func (h *UserHandler) Register(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{
			Code: http.StatusBadRequest,
			Msg:  "参数错误",
			Data: nil,
		})
		return
	}

	createdUser, err := h.userService.Register(req.Username, req.Email, req.Password)
	if err != nil {
		switch err {
		case service.ErrUsernameExists:
			c.JSON(http.StatusBadRequest, dto.Response{
				Code: http.StatusBadRequest,
				Msg:  "用户名被占用",
				Data: nil,
			})
		case service.ErrEmailExists:
			c.JSON(http.StatusBadRequest, dto.Response{
				Code: http.StatusBadRequest,
				Msg:  "邮箱被占用",
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
		Msg:  "注册成功",
		Data: dto.UserResponse{
			UserID: createdUser.ID,
		},
	})
}

// Login
// @Summary Log in a user
// @Description Log in a user with email and password
// @Tags users
// @Accept  json
// @Produce  json
// @Param   user     body    dto.LoginRequest     true        "User login info"
// @Success 200 {object} dto.LoginResponse "Login successful"
// @Failure 400 {object} dto.ErrorResponse "Invalid request body"
// @Failure 401 {object} dto.ErrorResponse "Unauthorized"
// @Router /login [post]
func (h *UserHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{
			Code: http.StatusBadRequest,
			Msg:  "参数错误",
			Data: nil,
		})
		return
	}

	token, err := h.userService.Login(req.Email, req.Password)
	if err != nil {
		switch err {
		case service.ErrInvalidCreds:
			c.JSON(http.StatusUnauthorized, dto.Response{
				Code: http.StatusUnauthorized,
				Msg:  "邮箱或密码错误",
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
		Msg:  "登录成功",
		Data: dto.LoginResponse{
			Token: token,
		},
	})
}
