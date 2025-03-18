package services

import (
	"errors"

	"github.com/plusone/models"
	"github.com/plusone/repositories"
	"github.com/plusone/utils"
	"gorm.io/gorm"
)

// UserService 用户服务层
type UserService struct {
	repo *repositories.UserRepository
	jwt  string
}

// NewUserService 创建用户服务实例
func NewUserService(repo *repositories.UserRepository, jwtSecret string) *UserService {
	return &UserService{
		repo: repo,
		jwt:  jwtSecret,
	}
}

// Register 用户注册
func (s *UserService) Register(username, password, email, nickname string) (*models.User, error) {
	// 检查用户名是否已存在
	_, err := s.repo.FindByUsername(username)
	if err == nil {
		return nil, errors.New("用户名已存在")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	// 创建新用户
	user := &models.User{
		Username: username,
		Email:    email,
		Nickname: nickname,
	}

	// 设置密码
	if err := user.SetPassword(password); err != nil {
		return nil, err
	}

	// 保存用户
	if err := s.repo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

// Login 用户登录
func (s *UserService) Login(username, password string) (string, error) {
	user, err := s.repo.FindByUsername(username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", errors.New("用户名不存在")
		}
		return "", err
	}

	// 验证密码
	if !user.CheckPassword(password) {
		return "", errors.New("密码错误")
	}

	// 生成JWT令牌
	token, err := utils.GenerateToken(user.ID, s.jwt)
	if err != nil {
		return "", err
	}

	return token, nil
}

// GetUserByID 通过ID获取用户
func (s *UserService) GetUserByID(id uint) (*models.User, error) {
	return s.repo.FindByID(id)
}
