package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/plusone/di"
	"github.com/plusone/middlewares"
)

// SetupRouter 配置路由
func SetupRouter(container *di.Container, jwtSecret string) *gin.Engine {
	r := gin.Default()

	userController := container.UserController

	// API组
	api := r.Group("/api")
	{
		// 公开路由
		api.POST("/register", userController.Register)
		api.POST("/login", userController.Login)

		// 需要认证的路由
		auth := api.Group("/user")
		auth.Use(middlewares.Auth(jwtSecret))
		{
			auth.GET("/info", userController.GetUserInfo)
		}
	}

	return r
}
