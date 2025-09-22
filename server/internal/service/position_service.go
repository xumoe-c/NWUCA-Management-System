package service

import (
	"NWUCA-Management-System/server/internal/model"
	"NWUCA-Management-System/server/internal/repository"
)

type PositionService interface {
	Create(name string) (*model.Position, error)
	GetAll() ([]model.Position, error)
	GetByID(id uint) (*model.Position, error)
	Update(id uint, name string) (*model.Position, error)
	Delete(id uint) error
}

type positionServiceImpl struct {
	repo repository.PositionRepository
}

func NewPositionService(repo repository.PositionRepository) PositionService {
	return &positionServiceImpl{repo: repo}
}

func (s *positionServiceImpl) Create(name string) (*model.Position, error) {
	position := &model.Position{Name: name}
	err := s.repo.Create(position)
	return position, err
}

func (s *positionServiceImpl) GetAll() ([]model.Position, error) {
	return s.repo.FindAll()
}

func (s *positionServiceImpl) GetByID(id uint) (*model.Position, error) {
	return s.repo.FindByID(id)
}

func (s *positionServiceImpl) Update(id uint, name string) (*model.Position, error) {
	position, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	position.Name = name
	err = s.repo.Update(position)
	return position, err
}

func (s *positionServiceImpl) Delete(id uint) error {
	return s.repo.Delete(id)
}
