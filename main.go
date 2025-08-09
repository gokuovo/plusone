package main

import (
	"log/slog"

	"github.com/plusone/config"
	"github.com/plusone/di"
	_ "github.com/plusone/docs" // 引入生成的 docs
	"github.com/plusone/models"
	"github.com/plusone/routes"
	"github.com/plusone/utils"
	"github.com/plusone/utils/logger"
)

// @title PlusOne API
// @version 1.0
// @description 这是一个使用 Go 语言编写的示例 Web 应用 API 文档
// @host localhost:8080
// @BasePath /api
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	// 初始化日志系统
	logger.Init()
	// 程序退出时关闭日志文件
	defer logger.CloseLogFile()

	// 加载配置
	cfg, err := config.LoadConfig()
	if err != nil {
		slog.Error("加载配置失败", "error", err)
		return
	}
	slog.Info("配置加载成功")

	// 连接数据库
	db, err := utils.ConnectDB(cfg.DBType, cfg.DBSource, cfg.DBLogLevel)
	if err != nil {
		slog.Error("连接数据库失败", "error", err)
		return
	}
	// 程序退出时关闭数据库连接
	defer utils.Close(db)
	slog.Info("数据库连接成功")

	// 连接 Redis
	redisClient, err := utils.ConnectRedis(cfg.RedisAddr, cfg.RedisPassword, cfg.RedisDB)
	if err != nil {
		slog.Error("连接 Redis 失败", "error", err)
		return
	}
	// 程序退出时关闭 Redis 连接
	defer redisClient.Close()
	slog.Info("Redis 连接成功")

	// 自动迁移表结构
	if err := db.AutoMigrate(&models.User{}); err != nil {
		slog.Error("数据库迁移失败", "error", err)
		return
	}
	slog.Info("数据库迁移完成")

	// 初始化依赖注入容器
	container := di.NewContainer(db, cfg.JWTSecret, redisClient)
	slog.Info("依赖注入容器初始化完成")

	// 设置路由
	router := routes.SetupRouter(container, cfg.JWTSecret)
	slog.Info("路由配置完成")

	// 启动服务器
	slog.Info("服务器启动", "port", cfg.ServerPort)
	if err := router.Run(":" + cfg.ServerPort); err != nil {
		slog.Error("服务器启动失败", "error", err)
		return
	}
}
