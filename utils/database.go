package utils

import (
	"fmt"
	"log"
	"log/slog"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// ConnectDB 连接到数据库
func ConnectDB(dbType, dbSource, dbLogLevel string) (*gorm.DB, error) {
	var db *gorm.DB
	var err error

	logLevel := parseLogLevel(dbLogLevel)

	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	}

	switch dbType {
	case "mysql":
		db, err = gorm.Open(mysql.Open(dbSource), gormConfig)
	case "sqlite":
		db, err = gorm.Open(sqlite.Open(dbSource), gormConfig)
	default:
		return nil, fmt.Errorf("不支持的数据库类型: %s", dbType)
	}

	if err != nil {
		return nil, err
	}

	log.Println("数据库连接成功")
	return db, nil
}

// Close 安全地关闭数据库连接
func Close(db *gorm.DB) {
	if db == nil {
		return
	}
	sqlDB, err := db.DB()
	if err != nil {
		slog.Error("获取底层数据库连接失败", "error", err)
		return
	}
	if err := sqlDB.Close(); err != nil {
		slog.Error("关闭数据库连接失败", "error", err)
	} else {
		slog.Info("数据库连接已成功关闭")
	}
}

// parseLogLevel 将字符串日志级别转换为 gorm 的 LogLevel
func parseLogLevel(level string) logger.LogLevel {
	switch strings.ToLower(level) {
	case "silent":
		return logger.Silent
	case "error":
		return logger.Error
	case "warn":
		return logger.Warn
	case "info":
		return logger.Info
	default:
		return logger.Info // 默认为 Info
	}
}
