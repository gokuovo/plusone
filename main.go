package main

import (
	"log"

	"github.com/oneplusone/config"
	"github.com/oneplusone/models"
	"github.com/oneplusone/routes"
	"github.com/oneplusone/utils"
	"github.com/oneplusone/di"
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

	// 初始化依赖注入容器
	container := di.NewContainer(db, cfg.JWTSecret)

	// 设置路由
	router := routes.SetupRouter(container, cfg.JWTSecret)

	// 启动服务器
	log.Printf("服务器已启动，监听端口 %s", cfg.ServerPort)
	router.Run(":" + cfg.ServerPort)
}
