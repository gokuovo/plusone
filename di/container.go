package di

import (
	"github.com/plusone/controllers"
	"github.com/plusone/repositories"
	"github.com/plusone/services"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// Container 依赖注入容器
type Container struct {
	UserController *controllers.UserController
}

// NewContainer 创建一个新的依赖注入容器
func NewContainer(db *gorm.DB, jwtSecret string, rdb *redis.Client) *Container {
	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository, jwtSecret, rdb)
	userController := controllers.NewUserController(userService)

	return &Container{
		UserController: userController,
	}
}
