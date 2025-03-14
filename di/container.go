package di

import (
	"github.com/oneplusone/controllers"
	"github.com/oneplusone/repositories"
	"github.com/oneplusone/services"
	"github.com/oneplusone/utils"
)

type Container struct {
	UserController *controllers.UserController
}

func NewContainer(db *utils.Database, jwtSecret string) *Container {
	// 初始化仓库
	userRepo := repositories.NewUserRepository(db)

	// 初始化服务
	userService := services.NewUserService(userRepo, jwtSecret)

	// 初始化控制器
	userController := controllers.NewUserController(userService)

	return &Container{
		UserController: userController,
	}
}