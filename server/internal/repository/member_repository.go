package repository

import (
	"NWUCA-Management-System/server/internal/model"

	"gorm.io/gorm"
)

type MemberRepository interface {
	Create(member *model.Member) error
	FindAll() ([]model.Member, error)
	FindByID(id uint) (*model.Member, error)
	Update(member *model.Member) error
	Delete(id uint) error
	FindByUserID(userID uint) (*model.Member, error)
}

type memberGormRepository struct {
	db *gorm.DB
}

func NewMemberRepository(db *gorm.DB) MemberRepository {
	return &memberGormRepository{db: db}
}

func (r *memberGormRepository) Create(member *model.Member) error {
	return r.db.Create(member).Error
}

func (r *memberGormRepository) FindAll() ([]model.Member, error) {
	var members []model.Member
	// Preload("User") 会在查询成员时，自动带上关联的 User 信息
	err := r.db.Preload("User").Find(&members).Error
	return members, err
}

func (r *memberGormRepository) FindByID(id uint) (*model.Member, error) {
	var member model.Member
	err := r.db.Preload("User").First(&member, id).Error
	return &member, err
}

func (r *memberGormRepository) Update(member *model.Member) error {
	return r.db.Save(member).Error
}

func (r *memberGormRepository) Delete(id uint) error {
	return r.db.Delete(&model.Member{}, id).Error
}

func (r *memberGormRepository) FindByUserID(userID uint) (*model.Member, error) {
	var member model.Member
	err := r.db.Where("user_id = ?", userID).First(&member).Error
	return &member, err
}
