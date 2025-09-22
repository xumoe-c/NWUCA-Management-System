package service

import (
	"NWUCA-Management-System/server/internal/model"
	"NWUCA-Management-System/server/internal/repository"
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type MemberService interface {
	CreateMember(req CreateMemberRequest) (*model.Member, error)
	GetAllMembers() ([]model.Member, error)
	UpdateMember(id uint, req UpdateMemberRequest) (*model.Member, error)
	DeleteMember(id uint) error
}

type CreateMemberRequest struct {
	Username  string `json:"username" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Name      string `json:"name" binding:"required"`
	StudentID string `json:"student_id" binding:"required"`
	Grade     string `json:"grade"`
	College   string `json:"college"`
	Major     string `json:"major"`
	Phone     string `json:"phone"`
}

type UpdateMemberRequest struct {
	Name      string `json:"name"`
	StudentID string `json:"student_id"`
	Grade     string `json:"grade"`
	College   string `json:"college"`
	Major     string `json:"major"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
}

type memberServiceImpl struct {
	db         *gorm.DB
	memberRepo repository.MemberRepository
	userRepo   repository.UserRepository
}

func NewMemberService(db *gorm.DB, memberRepo repository.MemberRepository, userRepo repository.UserRepository) MemberService {
	return &memberServiceImpl{
		db:         db,
		memberRepo: memberRepo,
		userRepo:   userRepo,
	}
}

func (s *memberServiceImpl) CreateMember(req CreateMemberRequest) (*model.Member, error) {
	// 使用事务确保用户和成员信息同时创建成功
	tx := s.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	// 检查用户名或邮箱是否已存在
	_, err := s.userRepo.FindByUsername(req.Username)
	if err == nil {
		tx.Rollback()
		return nil, errors.New("username already exists")
	}
	_, err = s.userRepo.FindByEmail(req.Email)
	if err == nil {
		tx.Rollback()
		return nil, errors.New("email already exists")
	}

	// 创建 User
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	user := &model.User{
		Username:     req.Username,
		Email:        req.Email,
		PasswordHash: string(hashedPassword),
		Role:         "member", // 新创建的会员默认为 'member' 角色
	}
	if err := tx.Create(user).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// 创建 Member
	member := &model.Member{
		UserID:    user.ID,
		Name:      req.Name,
		StudentID: req.StudentID,
		Grade:     req.Grade,
		College:   req.College,
		Major:     req.Major,
		Phone:     req.Phone,
		Email:     req.Email,
	}
	if err := tx.Create(member).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return member, tx.Commit().Error
}

func (s *memberServiceImpl) GetAllMembers() ([]model.Member, error) {
	return s.memberRepo.FindAll()
}

func (s *memberServiceImpl) UpdateMember(id uint, req UpdateMemberRequest) (*model.Member, error) {
	member, err := s.memberRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("member not found")
	}

	// 更新 Member 表信息
	if req.Name != "" {
		member.Name = req.Name
	}
	if req.StudentID != "" {
		member.StudentID = req.StudentID
	}
	if req.Grade != "" {
		member.Grade = req.Grade
	}
	if req.College != "" {
		member.College = req.College
	}
	if req.Major != "" {
		member.Major = req.Major
	}
	if req.Phone != "" {
		member.Phone = req.Phone
	}
	if req.Email != "" {
		member.Email = req.Email
	}

	if err := s.memberRepo.Update(member); err != nil {
		return nil, err
	}

	return member, nil
}

func (s *memberServiceImpl) DeleteMember(id uint) error {
	// 在事务中同时删除 member 和 user
	tx := s.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	member, err := s.memberRepo.FindByID(id)
	if err != nil {
		tx.Rollback()
		return errors.New("member not found")
	}

	// 删除 member
	if err := tx.Delete(&model.Member{}, id).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 删除 user
	if err := tx.Delete(&model.User{}, member.UserID).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
