package middlewares

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/plusone/response"
	"github.com/plusone/utils"
)

// Auth 认证中间件
func Auth(jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response.Error(c, errors.New("未提供认证令牌"))
			c.Abort()
			return
		}

		// 检查Bearer前缀
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			response.Error(c, errors.New("认证格式无效"))
			c.Abort()
			return
		}

		// 验证令牌
		tokenString := parts[1]
		claims, err := utils.ValidateToken(tokenString, jwtSecret)
		if err != nil {
			response.Error(c, errors.New("无效的令牌"))
			c.Abort()
			return
		}

		// 将用户ID存储在上下文中
		c.Set("userID", claims.UserID)
		c.Next()
	}
}
