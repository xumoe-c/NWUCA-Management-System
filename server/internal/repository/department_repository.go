package repository

import (
	"NWUCA-Management-System/server/internal/model"

	"gorm.io/gorm"
)

type DepartmentRepository interface {
	Create(department *model.Department) error
	FindAll() ([]model.Department, error)
	FindByID(id uint) (*model.Department, error)
	Update(department *model.Department) error
	Delete(id uint) error
}

type departmentGormRepository struct {
	db *gorm.DB
}

func NewDepartmentRepository(db *gorm.DB) DepartmentRepository {
	return &departmentGormRepository{db: db}
}

func (r *departmentGormRepository) Create(department *model.Department) error {
	return r.db.Create(department).Error
}

func (r *departmentGormRepository) FindAll() ([]model.Department, error) {
	var departments []model.Department
	err := r.db.Find(&departments).Error
	return departments, err
}

func (r *departmentGormRepository) FindByID(id uint) (*model.Department, error) {
	var department model.Department
	err := r.db.First(&department, id).Error
	return &department, err
}

func (r *departmentGormRepository) Update(department *model.Department) error {
	return r.db.Save(department).Error
}

func (r *departmentGormRepository) Delete(id uint) error {
	return r.db.Delete(&model.Department{}, id).Error
}
