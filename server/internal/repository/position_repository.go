package repository

import (
	"NWUCA-Management-System/server/internal/model"

	"gorm.io/gorm"
)

type PositionRepository interface {
	Create(position *model.Position) error
	FindAll() ([]model.Position, error)
	FindByID(id uint) (*model.Position, error)
	Update(position *model.Position) error
	Delete(id uint) error
}

type positionGormRepository struct {
	db *gorm.DB
}

func NewPositionRepository(db *gorm.DB) PositionRepository {
	return &positionGormRepository{db: db}
}

func (r *positionGormRepository) Create(position *model.Position) error {
	return r.db.Create(position).Error
}

func (r *positionGormRepository) FindAll() ([]model.Position, error) {
	var positions []model.Position
	err := r.db.Find(&positions).Error
	return positions, err
}

func (r *positionGormRepository) FindByID(id uint) (*model.Position, error) {
	var position model.Position
	err := r.db.First(&position, id).Error
	return &position, err
}

func (r *positionGormRepository) Update(position *model.Position) error {
	return r.db.Save(position).Error
}

func (r *positionGormRepository) Delete(id uint) error {
	return r.db.Delete(&model.Position{}, id).Error
}
