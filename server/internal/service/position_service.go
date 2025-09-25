package service

import (
	"NWUCA-Management-System/server/internal/model"
	"NWUCA-Management-System/server/internal/repository"
	"errors"
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

var (
	ErrNameExists        = errors.New("name already exists")
	ErrPositionNotExists = errors.New("position does not exists")
)

func NewPositionService(repo repository.PositionRepository) PositionService {
	return &positionServiceImpl{repo: repo}
}

func (s *positionServiceImpl) Create(name string) (*model.Position, error) {
	// 1. 检查name是否存在
	_, err := s.repo.FindByName(name)
	if err == nil {
		return nil, ErrNameExists
	}

	// 2. 创建position
	position := &model.Position{Name: name}
	err = s.repo.Create(position)
	if err != nil {
		return nil, err
	}
	return position, nil
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
		return nil, ErrPositionNotExists
	}

	position.Name = name
	err = s.repo.Update(position)
	if err == nil {
		return nil, err
	}
	return position, nil
}

func (s *positionServiceImpl) Delete(id uint) error {
	return s.repo.Delete(id)
}
