package controllers

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/plusone/dto"
	"github.com/plusone/response"
	"github.com/plusone/services"
	"github.com/plusone/utils/logger"
)

// UserController 用户控制器
type UserController struct {
	userService *services.UserService
}

// NewUserController 创建用户控制器实例
func NewUserController(userService *services.UserService) *UserController {
	return &UserController{userService: userService}
}

// Register
// @Summary 用户注册
// @Description 创建一个新用户
// @Tags Users
// @Accept json
// @Produce json
// @Param user body dto.RegisterInput true "用户信息"
// @Success 200 {object} response.Response{data=dto.UserOutput} "注册成功"
// @Failure 500 {object} response.Response "服务器错误"
// @Router /register [post]
func (c *UserController) Register(ctx *gin.Context) {
	var input dto.RegisterInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		logger.CtxErrorf(ctx, "参数绑定失败: %v", err)
		response.Error(ctx, err)
		return
	}

	user, err := c.userService.Register(ctx, input.Username, input.Password, input.Email, input.Nickname)
	if err != nil {
		logger.CtxErrorf(ctx, "用户注册失败: %v", err)
		response.Error(ctx, err)
		return
	}

	logger.CtxInfof(ctx, "用户注册成功: %s", user.Username)
	response.Success(ctx, dto.NewUserOutput(user))
}

// Login
// @Summary 用户登录
// @Description 用户使用用户名和密码登录，获取JWT
// @Tags Users
// @Accept json
// @Produce json
// @Param credentials body dto.LoginInput true "登录凭证"
// @Success 200 {object} response.Response{data=dto.LoginOutput} "登录成功"
// @Failure 500 {object} response.Response "服务器错误"
// @Router /login [post]
func (c *UserController) Login(ctx *gin.Context) {
	var input dto.LoginInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		logger.CtxErrorf(ctx, "参数绑定失败: %v", err)
		response.Error(ctx, err)
		return
	}

	token, err := c.userService.Login(ctx, input.Username, input.Password)
	if err != nil {
		logger.CtxErrorf(ctx, "用户登录失败: %v", err)
		response.Error(ctx, err)
		return
	}

	logger.CtxInfof(ctx, "用户登录成功: %s", input.Username)
	response.Success(ctx, dto.LoginOutput{Token: token})
}

// GetUserInfo
// @Summary 获取当前用户信息
// @Description 根据JWT获取当前登录用户的信息
// @Tags Users
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} response.Response{data=dto.UserOutput} "获取成功"
// @Failure 500 {object} response.Response "服务器错误"
// @Router /user/info [get]
func (c *UserController) GetUserInfo(ctx *gin.Context) {
	userID, exists := ctx.Get("userID")
	if !exists {
		err := errors.New("未认证: 无法从上下文中获取用户ID")
		logger.CtxErrorf(ctx, err.Error())
		response.Error(ctx, err)
		return
	}

	user, err := c.userService.GetUserByID(ctx, userID.(uint))
	if err != nil {
		logger.CtxErrorf(ctx, "获取用户信息失败, userID: %d, error: %v", userID, err)
		response.Error(ctx, err)
		return
	}

	logger.CtxInfof(ctx, "获取用户信息成功, userID: %d", userID)
	response.Success(ctx, dto.NewUserOutput(user))
}
