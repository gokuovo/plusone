package utils

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Database 数据库连接管理
type Database struct {
	*gorm.DB
}

// ConnectDB 连接到数据库
func ConnectDB(dbType string, dbSource string) (*Database, error) {
	var db *gorm.DB
	var err error

	switch dbType {
	case "mysql":
		db, err = gorm.Open(mysql.Open(dbSource), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
	case "sqlite":
		db, err = gorm.Open(sqlite.Open(dbSource), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
	default:
		return nil, fmt.Errorf("不支持的数据库类型: %s", dbType)
	}

	if err != nil {
		return nil, err
	}

	log.Println("数据库连接成功")
	return &Database{db}, nil
}
