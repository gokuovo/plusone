package main

import (
	"log"

	"github.com/oneplusone/config"
	"github.com/oneplusone/controllers"
	"github.com/oneplusone/models"
	"github.com/oneplusone/repositories"
	"github.com/oneplusone/routes"
	"github.com/oneplusone/services"
	"github.com/oneplusone/utils"
)

func main() {
	// 加载配置
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("加载配置失败: %v", err)
	}

	// 连接数据库
	db, err := utils.ConnectDB(cfg.DBType, cfg.DBSource)
	if err != nil {
		log.Fatalf("连接数据库失败: %v", err)
	}

	// 自动迁移表结构
	db.AutoMigrate(&models.User{})

	// 初始化仓库、服务和控制器
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo, cfg.JWTSecret)
	userController := controllers.NewUserController(userService)

	// 设置路由
	router := routes.SetupRouter(userController, cfg.JWTSecret)

	// 启动服务器
	log.Printf("服务器已启动，监听端口 %s", cfg.ServerPort)
	router.Run(":" + cfg.ServerPort)
}
