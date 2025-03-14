package config

import (
	"os"

	"github.com/joho/godotenv"
)

// Config 应用配置
type Config struct {
	DBType     string
	DBSource   string
	JWTSecret  string
	ServerPort string
}

// LoadConfig 从环境变量加载配置
func LoadConfig() (*Config, error) {
	// 加载.env文件
	_ = godotenv.Load()

	return &Config{
		DBType:     getEnv("DB_TYPE", "sqlite"), // mysql 或 sqlite
		DBSource:   getEnv("DB_SOURCE", "oneplusone.db"),
		JWTSecret:  getEnv("JWT_SECRET", "your_jwt_secret_key"),
		ServerPort: getEnv("SERVER_PORT", "8080"),
	}, nil
}

// getEnv 获取环境变量，如果不存在则返回默认值
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
