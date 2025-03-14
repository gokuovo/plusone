package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User 用户模型
type User struct {
	gorm.Model
	Username string `gorm:"size:50;not null;uniqueIndex" json:"username"`
	Password string `gorm:"size:100;not null" json:"-"` // 不在JSON中显示密码
	Email    string `gorm:"size:100;uniqueIndex" json:"email"`
	Nickname string `gorm:"size:50" json:"nickname"`
}

// SetPassword 设置加密后的密码
func (u *User) SetPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

// CheckPassword 检查密码是否正确
func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
