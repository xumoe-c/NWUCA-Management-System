package dto

import "time"

// CreateMemberRequest 定义了创建成员的请求结构
type CreateMemberRequest struct {
	Name         string    `json:"name" binding:"required"`
	JoinDate     time.Time `json:"join_date" binding:"required"`
	DepartmentID uint      `json:"department_id" binding:"required"`
	PositionID   uint      `json:"position_id" binding:"required"`
	Email        string    `json:"email" binding:"required,email"`
	PhoneNumber  string    `json:"phone_number"`
}

// UpdateMemberRequest 定义了更新成员的请求结构
type UpdateMemberRequest struct {
	Name         *string    `json:"name"`
	JoinDate     *time.Time `json:"join_date"`
	DepartmentID *uint      `json:"department_id"`
	PositionID   *uint      `json:"position_id"`
	Email        *string    `json:"email,email"`
	PhoneNumber  *string    `json:"phone_number"`
}
