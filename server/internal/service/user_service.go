package service

import (
	"NWUCA-Management-System/server/internal/errors"
	"NWUCA-Management-System/server/internal/model"
	"NWUCA-Management-System/server/internal/repository"
	"NWUCA-Management-System/server/internal/util/auth"
	"errors"

	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// UserService 接口定义了用户服务
type UserService interface {
	Register(username, email, password string) (*model.User, error)
	Login(email, password string) (string, error)
}

// userServiceImpl 是 UserService 的实现
type userServiceImpl struct {
	userRepo   repository.UserRepository // 依赖仓库层的接口
	jwtSecret  string
	jwtExpDays int
}

// NewUserService 创建一个新的 UserService 实例
func NewUserService(userRepo repository.UserRepository, jwtSecret string, jwtExpDays int) UserService {
	return &userServiceImpl{
		userRepo:   userRepo,
		jwtSecret:  jwtSecret,
		jwtExpDays: jwtExpDays,
	}
}

// Register 实现了用户注册的业务逻辑
func (s *userServiceImpl) Register(username, email, password string) (*model.User, error) {
	// 1. 检查用户是否已存在
	_, err := s.userRepo.FindByUsername(username)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	} else {
		return nil, apperrors.ErrUsernameExists
	}

	_, err = s.userRepo.FindByEmail(email)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	} else {
		return nil, apperrors.ErrEmailExists
	}

	// 2. 哈希密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// 3. 创建用户
	newUser := &model.User{
		Username:     username,
		PasswordHash: string(hashedPassword),
		Email:        email,
		Role:         "member", // 默认角色为 'member'
	}

	// 4. 保存到数据库
	err = s.userRepo.Create(newUser)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}

// Login 实现了用户登录的业务逻辑
func (s *userServiceImpl) Login(email, password string) (string, error) {
	// WARNING: 调试环境使用，生产环境请删除！
	if email == viper.GetString("admin_email") {
		if password == viper.GetString("admin_password") {
			return auth.GenerateToken(1, "admin", s.jwtSecret, s.jwtExpDays)
		}
	}

	// 1. 根据邮箱查找用户
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return "", apperrors.ErrInvalidCredits
	}

	// 2. 验证密码
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return "", apperrors.ErrInvalidCredits
	}

	// 3. 生成JWT
	token, err := auth.GenerateToken(user.ID, user.Role, s.jwtSecret, s.jwtExpDays)
	if err != nil {
		return "", err
	}

	return token, nil
}
