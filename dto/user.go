package dto

import "github.com/plusone/models"

// RegisterInput 用户注册的输入
type RegisterInput struct {
	Username string `json:"username" binding:"required" example:"testuser"`
	Password string `json:"password" binding:"required" example:"password123"`
	Email    string `json:"email" binding:"required,email" example:"test@example.com"`
	Nickname string `json:"nickname" example:"Tester"`
}

// LoginInput 用户登录的输入
type LoginInput struct {
	Username string `json:"username" binding:"required" example:"testuser"`
	Password string `json:"password" binding:"required" example:"password123"`
}

// LoginOutput 用户登录的输出
type LoginOutput struct {
	Token string `json:"token"`
}

// UserOutput 用户信息的标准输出
type UserOutput struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Nickname string `json:"nickname"`
}

// NewUserOutput 将 models.User 转换为 UserOutput DTO
func NewUserOutput(user *models.User) UserOutput {
	return UserOutput{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Nickname: user.Nickname,
	}
}
