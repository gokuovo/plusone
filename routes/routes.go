package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/plusone/di"
	"github.com/plusone/middlewares"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// SetupRouter 配置路由
func SetupRouter(container *di.Container, jwtSecret string) *gin.Engine {
	// 使用 gin.New() 创建一个不带默认中间件的引擎
	r := gin.New()

	// 全局中间件
	// 1. 日志中间件
	r.Use(middlewares.LoggerMiddleware())
	// 2. 恢复中间件，防止 panic 导致程序崩溃
	r.Use(gin.Recovery())

	// 添加 Swagger 路由
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
