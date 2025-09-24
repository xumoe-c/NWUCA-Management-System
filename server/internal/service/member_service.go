package service

import (
	"NWUCA-Management-System/server/internal/dto"
	"NWUCA-Management-System/server/internal/model"
	"NWUCA-Management-System/server/internal/repository"
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type MemberService interface {
	CreateMember(req dto.CreateMemberRequest) (*model.Member, error)
	GetAllMembers() ([]model.Member, error)
	UpdateMember(id uint, req dto.UpdateMemberRequest) (*model.Member, error)
	DeleteMember(id uint) error
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

func (s *memberServiceImpl) CreateMember(req dto.CreateMemberRequest) (*model.Member, error) {
	// 使用事务确保用户和成员信息同时创建成功
	tx := s.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	// 检查邮箱是否已存在
	_, err := s.userRepo.FindByEmail(req.Email)
	if err == nil {
		tx.Rollback()
		return nil, errors.New("email already exists")
	}

	// 创建 User
	// 注意：在实际生产中，密码应该由前端传递，这里为了简化，暂时设为默认值
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("default_password"), bcrypt.DefaultCost)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	user := &model.User{
		Username:     req.Name, // 暂时使用成员姓名作为用户名
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
		UserID:       user.ID,
		Name:         req.Name,
		JoinDate:     req.JoinDate,
		DepartmentID: req.DepartmentID,
		PositionID:   req.PositionID,
		Email:        req.Email,
		PhoneNumber:  req.PhoneNumber,
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

func (s *memberServiceImpl) UpdateMember(id uint, req dto.UpdateMemberRequest) (*model.Member, error) {
	member, err := s.memberRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("member not found")
	}

	// 更新 Member 表信息
	if req.Name != nil {
		member.Name = *req.Name
	}
	if req.JoinDate != nil {
		member.JoinDate = *req.JoinDate
	}
	if req.DepartmentID != nil {
		member.DepartmentID = *req.DepartmentID
	}
	if req.PositionID != nil {
		member.PositionID = *req.PositionID
	}
	if req.Email != nil {
		member.Email = *req.Email
	}
	if req.PhoneNumber != nil {
		member.PhoneNumber = *req.PhoneNumber
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
