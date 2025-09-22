package model

import (
	"time"

	"gorm.io/gorm"
)

// 1. 用户表 (Users)
// 存储系统登录账号信息
type User struct {
	gorm.Model
	Username     string `gorm:"size:255;not null;unique" json:"username"`
	Email        string `gorm:"size:255;not null;unique" json:"email"`
	PasswordHash string `gorm:"size:255;not null" json:"-"`                    // 存储哈希后的密码，json:"-"表示不通过API暴露
	Role         string `gorm:"size:50;not null;default:'member'" json:"role"` // 角色 (e.g., 'admin', 'member')
}

// 2. 会员表 (Members)
// 存储协会成员的详细档案
type Member struct {
	gorm.Model
	UserID                 uint   `gorm:"not null;unique"`
	User                   User   `gorm:"foreignKey:UserID"` // 关联用户表
	Name                   string `gorm:"size:255;not null"`
	StudentID              string `gorm:"size:50;not null;unique"`
	Grade                  string `gorm:"size:50"` // 年级，如 "2022级"
	College                string `gorm:"size:100"`
	Major                  string `gorm:"size:100"`
	Phone                  string `gorm:"size:20;unique"`
	Email                  string `gorm:"size:255;not null;unique"`
	JoinDate               time.Time
	ExpectedGraduationDate time.Time
	Status                 string       `gorm:"size:50;not null;default:'active'"` // 状态: active, graduated, inactive
	AvatarURL              string       `gorm:"size:255"`
	Assignments            []Assignment // 一个会员可以有多个任期分配
}

// 3. 指导老师表 (Advisors)
// 存储指导老师的档案
type Advisor struct {
	gorm.Model
	UserID       uint   `gorm:"not null;unique"`
	User         User   `gorm:"foreignKey:UserID"`
	Name         string `gorm:"size:255;not null"`
	College      string `gorm:"size:100"`
	Title        string `gorm:"size:100"` // 职称
	Phone        string `gorm:"size:20;unique"`
	Email        string `gorm:"size:255;not null;unique"`
	ResearchArea string `gorm:"size:255"` // 研究方向
}

// 4. 部门表 (Departments)
// 定义协会的组织单元
type Department struct {
	gorm.Model
	Name        string       `gorm:"size:100;not null;unique"`
	Description string       `gorm:"type:text"`
	ParentID    *uint        // 支持层级结构 (父部门ID)，指针类型以允许为NULL
	Parent      *Department  `gorm:"foreignKey:ParentID"`
	Children    []Department `gorm:"foreignKey:ParentID"`
}

// 5. 职务表 (Positions)
// 定义协会内的职务
type Position struct {
	gorm.Model
	Name        string `gorm:"size:100;not null;unique"` // 如 "会长", "部长", "部员"
	AccessLevel int    `gorm:"not null;default:1"`       // 权限等级，数字越大权限越高
}

// 6. 任期分配表 (Assignments)
// 核心关联表，将会员、部门、职务和任期关联起来
type Assignment struct {
	gorm.Model
	MemberID     uint       `gorm:"not null"`
	Member       Member     `gorm:"foreignKey:MemberID"`
	DepartmentID uint       `gorm:"not null"`
	Department   Department `gorm:"foreignKey:DepartmentID"`
	PositionID   uint       `gorm:"not null"`
	Position     Position   `gorm:"foreignKey:PositionID"`
	StartDate    time.Time  `gorm:"not null"`
	EndDate      time.Time  `gorm:"not null"`
	Status       string     `gorm:"size:50;not null;default:'active'"` // 状态: active, ended
}
