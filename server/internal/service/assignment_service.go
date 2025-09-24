package service

import (
	"NWUCA-Management-System/server/internal/dto"
	"NWUCA-Management-System/server/internal/model"
	"NWUCA-Management-System/server/internal/repository"
)

type AssignmentService interface {
	CreateAssignment(req dto.CreateAssignmentRequest) (*model.Assignment, error)
	GetAllAssignments() ([]model.Assignment, error)
	UpdateAssignment(id uint, req dto.UpdateAssignmentRequest) (*model.Assignment, error)
	DeleteAssignment(id uint) error
}

type assignmentServiceImpl struct {
	repo repository.AssignmentRepository
}

func NewAssignmentService(repo repository.AssignmentRepository) AssignmentService {
	return &assignmentServiceImpl{repo: repo}
}

func (s *assignmentServiceImpl) CreateAssignment(req dto.CreateAssignmentRequest) (*model.Assignment, error) {
	assignment := &model.Assignment{
		Title:       req.Title,
		Description: req.Description,
		CreatedBy:   req.CreatedBy,
		AssigneeID:  req.AssigneeID,
		Status:      "pending", // 默认状态
	}
	err := s.repo.Create(assignment)
	return assignment, err
}

func (s *assignmentServiceImpl) GetAllAssignments() ([]model.Assignment, error) {
	return s.repo.FindAll()
}

func (s *assignmentServiceImpl) UpdateAssignment(id uint, req dto.UpdateAssignmentRequest) (*model.Assignment, error) {
	assignment, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if req.Title != nil {
		assignment.Title = *req.Title
	}
	if req.Description != nil {
		assignment.Description = *req.Description
	}
	if req.Status != nil {
		assignment.Status = *req.Status
	}
	if req.AssigneeID != nil {
		assignment.AssigneeID = *req.AssigneeID
	}

	err = s.repo.Update(assignment)
	return assignment, err
}

func (s *assignmentServiceImpl) DeleteAssignment(id uint) error {
	return s.repo.Delete(id)
}
