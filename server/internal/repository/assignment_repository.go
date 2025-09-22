package repository

import (
	"NWUCA-Management-System/server/internal/model"

	"gorm.io/gorm"
)

type AssignmentRepository interface {
	Create(assignment *model.Assignment) error
	FindAll() ([]model.Assignment, error)
	FindByID(id uint) (*model.Assignment, error)
	Update(assignment *model.Assignment) error
	Delete(id uint) error
}

type assignmentGormRepository struct {
	db *gorm.DB
}

func NewAssignmentRepository(db *gorm.DB) AssignmentRepository {
	return &assignmentGormRepository{db: db}
}

func (r *assignmentGormRepository) Create(assignment *model.Assignment) error {
	return r.db.Create(assignment).Error
}

func (r *assignmentGormRepository) FindAll() ([]model.Assignment, error) {
	var assignments []model.Assignment
	// 预加载关联数据
	err := r.db.Preload("Member").Preload("Department").Preload("Position").Find(&assignments).Error
	return assignments, err
}

func (r *assignmentGormRepository) FindByID(id uint) (*model.Assignment, error) {
	var assignment model.Assignment
	err := r.db.Preload("Member").Preload("Department").Preload("Position").First(&assignment, id).Error
	return &assignment, err
}

func (r *assignmentGormRepository) Update(assignment *model.Assignment) error {
	return r.db.Save(assignment).Error
}

func (r *assignmentGormRepository) Delete(id uint) error {
	return r.db.Delete(&model.Assignment{}, id).Error
}
