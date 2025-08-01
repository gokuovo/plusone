package repositories

import (
	"context"

	"github.com/plusone/models"
	"github.com/plusone/utils"
)

// UserRepository 用户数据访问层
type UserRepository struct {
	db *utils.Database
}

// NewUserRepository 创建用户仓库实例
func NewUserRepository(db *utils.Database) *UserRepository {
	return &UserRepository{db: db}
}

// Create 创建新用户
func (r *UserRepository) Create(ctx context.Context, user *models.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}

// FindByID 通过ID查找用户
func (r *UserRepository) FindByID(ctx context.Context, id uint) (*models.User, error) {
	var user models.User
	err := r.db.WithContext(ctx).First(&user, id).Error
	return &user, err
}

// FindByUsername 通过用户名查找用户
func (r *UserRepository) FindByUsername(ctx context.Context, username string) (*models.User, error) {
	var user models.User
	err := r.db.WithContext(ctx).Where("username = ?", username).First(&user).Error
	return &user, err
}

// Update 更新用户信息
func (r *UserRepository) Update(ctx context.Context, user *models.User) error {
	return r.db.WithContext(ctx).Save(user).Error
}
