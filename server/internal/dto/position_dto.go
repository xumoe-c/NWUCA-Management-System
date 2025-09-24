package dto

// CreatePositionRequest 定义了创建职位的请求结构
type CreatePositionRequest struct {
	Name string `json:"name" binding:"required"`
}

// UpdatePositionRequest 定义了更新职位的请求结构
type UpdatePositionRequest struct {
	Name string `json:"name" binding:"required"`
}
