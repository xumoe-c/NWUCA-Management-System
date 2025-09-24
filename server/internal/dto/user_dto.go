package dto

// RegisterRequest 定义了注册请求的 JSON 结构
type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

// LoginRequest 定义了登录请求的 JSON 结构
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// UserResponse 定义了成功注册用户的响应结构
type UserResponse struct {
	UserID uint `json:"user_id"`
}

// LoginResponse 定义了成功登录的响应结构
type LoginResponse struct {
	Token string `json:"token"`
}
