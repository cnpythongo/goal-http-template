package controller

import (
	"github.com/cnpythongo/goal/apps/account/model"
	"github.com/cnpythongo/goal/apps/account/service"
	"github.com/cnpythongo/goal/pkg/response"
	"github.com/gin-gonic/gin"
)

type IUserController interface {
	// 创建用户
	CreateUser(c *gin.Context)
	// 根据ID获取用户
	GetUserById(c *gin.Context)
	// 根据UUID获取用户
	GetUserByUuid(c *gin.Context)
	// 获取用户查询集
	GetUserList(c *gin.Context)
}

type UserController struct {
	UserSvc service.IUserService `inject:"UserSvc"`
}

func (u *UserController) CreateUser(c *gin.Context) {
	var payload model.User
	err := c.ShouldBindJSON(&payload)
	if err != nil {
		response.FailJsonResp(c, "提交表单数据不正确")
		return
	}
}

func (u *UserController) GetUserById(c *gin.Context) {
	panic("implement me")
}

func (u *UserController) GetUserByUuid(c *gin.Context) {
	panic("implement me")
}

func (u *UserController) GetUserList(c *gin.Context) {
	panic("implement me")
}
