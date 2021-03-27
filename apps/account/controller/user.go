package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	
	"github.com/cnpythongo/goal/apps/account/model"
	"github.com/cnpythongo/goal/apps/account/service"
	"github.com/cnpythongo/goal/pkg/response"
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
	payload := model.NewUser()
	err := c.ShouldBindJSON(payload)
	if err != nil {
		response.FailJsonResp(c, "提交表单数据不正确")
		return
	}
	user, err := u.UserSvc.CreateUser(payload)
	if err != nil {
		response.FailJsonResp(c, "创建用户失败")
		return
	}
	response.SuccessJsonResp(c, user, nil)
}

func (u *UserController) GetUserById(c *gin.Context) {
	panic("implement me")
}

func (u *UserController) GetUserByUuid(c *gin.Context) {
	uid := c.Param("uid")
	result, err := u.UserSvc.GetUserByUuid(uid)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			response.FailJsonResp(c, "用户不存在")
		} else {
			response.FailJsonResp(c, "查询用户失败")
		}
		return
	}
	response.SuccessJsonResp(c, result, nil)
}

func (u *UserController) GetUserList(c *gin.Context) {
	panic("implement me")
}
