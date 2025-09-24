package dto

// CreateAssignmentRequest 定义了创建任务的请求结构
type CreateAssignmentRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	CreatedBy   uint   `json:"created_by" binding:"required"`
	AssigneeID  uint   `json:"assignee_id"`
}

// UpdateAssignmentRequest 定义了更新任务的请求结构
type UpdateAssignmentRequest struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Status      *string `json:"status"`
	AssigneeID  *uint   `json:"assignee_id"`
}
