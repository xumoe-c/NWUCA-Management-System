package service

import (
	apperrors "NWUCA-Management-System/server/internal/errors"
	"NWUCA-Management-System/server/internal/model"
	"NWUCA-Management-System/server/internal/repository"
	"errors"

	"gorm.io/gorm"
)

type DepartmentService interface {
	Create(name string) (*model.Department, error)
	GetAll() ([]model.Department, error)
	GetByID(id uint) (*model.Department, error)
	Update(id uint, name string) (*model.Department, error)
	Delete(id uint) error
}

type departmentServiceImpl struct {
	repo repository.DepartmentRepository
}

func NewDepartmentService(repo repository.DepartmentRepository) DepartmentService {
	return &departmentServiceImpl{repo: repo}
}

func (s *departmentServiceImpl) Create(name string) (*model.Department, error) {
	_, err := s.repo.FindByName(name)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	} else {
		return nil, apperrors.ErrDepartmentExists
	}

	department := &model.Department{Name: name}
	err = s.repo.Create(department)
	return department, err
}

func (s *departmentServiceImpl) GetAll() ([]model.Department, error) {
	return s.repo.FindAll()
}

func (s *departmentServiceImpl) GetByID(id uint) (*model.Department, error) {
	return s.repo.FindByID(id)
}

func (s *departmentServiceImpl) Update(id uint, name string) (*model.Department, error) {
	department, err := s.repo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperrors.ErrNotFound
		} else {
			return nil, err
		}
	}

	department.Name = name
	err = s.repo.Update(department)
	if err != nil {
		return nil, err
	}
	return department, nil
}

func (s *departmentServiceImpl) Delete(id uint) error {
	return s.repo.Delete(id)
}
