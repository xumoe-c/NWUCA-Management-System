package dto

// CreateDepartmentRequest 定义了创建部门的请求结构
type CreateDepartmentRequest struct {
	Name string `json:"name" binding:"required"`
}

// UpdateDepartmentRequest 定义了更新部门的请求结构
type UpdateDepartmentRequest struct {
	Name string `json:"name" binding:"required"`
}
