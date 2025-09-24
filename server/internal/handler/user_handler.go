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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdUser, err := h.userService.Register(req.Username, req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, dto.UserResponse{
		Message: "User registered successfully",
		UserID:  createdUser.ID,
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.userService.Login(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.LoginResponse{Token: token})
}
