package service

import (
	"NWUCA-Management-System/server/internal/model"
	"NWUCA-Management-System/server/internal/repository"
	"time"
)

type AssignmentService interface {
	CreateAssignment(req CreateAssignmentRequest) (*model.Assignment, error)
	GetAllAssignments() ([]model.Assignment, error)
	UpdateAssignment(id uint, req UpdateAssignmentRequest) (*model.Assignment, error)
	DeleteAssignment(id uint) error
}

type CreateAssignmentRequest struct {
	MemberID     uint       `json:"member_id" binding:"required"`
	DepartmentID uint       `json:"department_id" binding:"required"`
	PositionID   uint       `json:"position_id" binding:"required"`
	StartDate    time.Time  `json:"start_date" binding:"required"`
	EndDate      *time.Time `json:"end_date"` // 使用指针以允许为空
}

type UpdateAssignmentRequest struct {
	DepartmentID uint       `json:"department_id"`
	PositionID   uint       `json:"position_id"`
	StartDate    time.Time  `json:"start_date"`
	EndDate      *time.Time `json:"end_date"`
}

type assignmentServiceImpl struct {
	repo repository.AssignmentRepository
}

func NewAssignmentService(repo repository.AssignmentRepository) AssignmentService {
	return &assignmentServiceImpl{repo: repo}
}

func (s *assignmentServiceImpl) CreateAssignment(req CreateAssignmentRequest) (*model.Assignment, error) {
	assignment := &model.Assignment{
		MemberID:     req.MemberID,
		DepartmentID: req.DepartmentID,
		PositionID:   req.PositionID,
		StartDate:    req.StartDate,
		EndDate:      *req.EndDate,
	}
	err := s.repo.Create(assignment)
	return assignment, err
}

func (s *assignmentServiceImpl) GetAllAssignments() ([]model.Assignment, error) {
	return s.repo.FindAll()
}

func (s *assignmentServiceImpl) UpdateAssignment(id uint, req UpdateAssignmentRequest) (*model.Assignment, error) {
	assignment, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if req.DepartmentID != 0 {
		assignment.DepartmentID = req.DepartmentID
	}
	if req.PositionID != 0 {
		assignment.PositionID = req.PositionID
	}
	if !req.StartDate.IsZero() {
		assignment.StartDate = req.StartDate
	}
	assignment.EndDate = *req.EndDate // 允许将 EndDate 更新为 nil

	err = s.repo.Update(assignment)
	return assignment, err
}

func (s *assignmentServiceImpl) DeleteAssignment(id uint) error {
	return s.repo.Delete(id)
}
