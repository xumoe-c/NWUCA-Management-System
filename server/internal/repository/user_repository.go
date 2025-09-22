package repository

import (
	"NWUCA-Management-System/server/internal/model"

	"gorm.io/gorm"
)

// UserRepository 接口定义了用户仓库需要实现的方法
type UserRepository interface {
	Create(user *model.User) error
	FindByUsername(username string) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
}

// userGormRepository 是 UserRepository 的 GORM 实现
type userGormRepository struct {
	db *gorm.DB
}

// NewUserRepository 创建一个 UserRepository 的 GORM 实现实例
// 这里是关键：构造函数返回的是接口，而不是具体的 struct
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userGormRepository{db: db}
}

// Create 实现创建用户的方法
func (r *userGormRepository) Create(user *model.User) error {
	return r.db.Create(user).Error
}

// FindByUsername 实现根据用户名查找用户的方法
func (r *userGormRepository) FindByUsername(username string) (*model.User, error) {
	var user model.User
	err := r.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// FindByEmail 根据邮箱查找用户
func (r *userGormRepository) FindByEmail(email string) (*model.User, error) {
	var user model.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
