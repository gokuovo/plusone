package services

import (
	"context"
	"errors"

	"github.com/plusone/models"
	"github.com/plusone/repositories"
	"github.com/plusone/utils"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// UserService 用户服务层
type UserService struct {
	repo *repositories.UserRepository
	jwt  string
	rdb  *redis.Client
}

// NewUserService 创建用户服务实例
func NewUserService(repo *repositories.UserRepository, jwtSecret string, rdb *redis.Client) *UserService {
	return &UserService{
		repo: repo,
		jwt:  jwtSecret,
		rdb:  rdb,
	}
}

// Register 用户注册
func (s *UserService) Register(ctx context.Context, username, password, email, nickname string) (*models.User, error) {
	var user *models.User
	// GORM 事务
	err := s.repo.Transaction(func(tx *gorm.DB) error {
		// 使用事务作用域的 repository
		txRepo := repositories.NewUserRepository(tx)

		// 1. 检查用户名是否已存在
		_, err := txRepo.FindByUsername(ctx, username)
		if err == nil {
			return errors.New("用户名已存在")
		}
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		// 2. 创建新用户
		newUser := &models.User{
			Username: username,
			Email:    email,
			Nickname: nickname,
		}

		// 3. 设置密码
		if err := newUser.SetPassword(password); err != nil {
			return err
		}

		// 4. 保存用户
		if err := txRepo.Create(ctx, newUser); err != nil {
			return err
		}

		user = newUser
		return nil // 事务提交
	})

	return user, err
}

// Login 用户登录
func (s *UserService) Login(ctx context.Context, username, password string) (string, error) {
	user, err := s.repo.FindByUsername(ctx, username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", errors.New("用户不存在")
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
func (s *UserService) GetUserByID(ctx context.Context, id uint) (*models.User, error) {
	user, err := s.repo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("用户不存在")
		}
		return nil, err
	}
	return user, nil
}
