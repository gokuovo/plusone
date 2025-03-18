package main

import (
	"github.com/plusone/config"
	"github.com/plusone/di"
	"github.com/plusone/models"
	"github.com/plusone/routes"
	"github.com/plusone/utils"
	"github.com/plusone/utils/logger"
)

func main() {
	// 程序退出时关闭日志文件
	defer logger.CloseLogFiles()

	// 加载配置
	cfg, err := config.LoadConfig()
	if err != nil {
		logger.Error("加载配置失败: %v", err)
		return
	}
	logger.Info("配置加载成功")

	// 连接数据库
	db, err := utils.ConnectDB(cfg.DBType, cfg.DBSource)
	if err != nil {
		logger.Error("连接数据库失败: %v", err)
		return
	}
	logger.Info("数据库连接成功")

	// 自动迁移表结构
	if err := db.AutoMigrate(&models.User{}); err != nil {
		logger.Error("数据库迁移失败: %v", err)
		return
	}
	logger.Info("数据库迁移完成")

	// 初始化依赖注入容器
	container := di.NewContainer(db, cfg.JWTSecret)
	logger.Info("依赖注入容器初始化完成")

	// 设置路由
	router := routes.SetupRouter(container, cfg.JWTSecret)
	logger.Info("路由配置完成")

	// 启动服务器
	logger.Info("服务器启动，监听端口 %s", cfg.ServerPort)
	if err := router.Run(":" + cfg.ServerPort); err != nil {
		logger.Error("服务器启动失败: %v", err)
		return
	}
}
