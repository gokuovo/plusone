package di

import (
	"github.com/plusone/controllers"
	"github.com/plusone/repositories"
	"github.com/plusone/services"
	"github.com/plusone/utils"
)

type Container struct {
	UserController *controllers.UserController
}

func NewContainer(db *utils.Database, jwtSecret string) *Container {
	// 初始化仓库
	userRepo := repositories.NewUserRepository(db)

	// 初始化服务, 注入 db 和 repo
	userService := services.NewUserService(db, userRepo, jwtSecret)

	// 初始化控制器
	userController := controllers.NewUserController(userService)

	return &Container{
		UserController: userController,
	}
}
