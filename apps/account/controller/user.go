package controller

import (
	"github.com/cnpythongo/goal/pkg/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"

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
	eu, _ := u.UserSvc.GetUserByUsername(payload.Username)
	if eu != nil {
		response.FailJsonResp(c, "用户名已存在，请换一个")
		return
	}
	ue, _ := u.UserSvc.GetUserByEmail(payload.Email)
	if ue != nil {
		response.FailJsonResp(c, "邮箱地址已存在，请换一个")
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
	pk := c.Param("id")
	id, e := strconv.Atoi(pk)
	if e != nil {
		response.FailJsonResp(c, "用户id不正确")
		return
	}
	result, err := u.UserSvc.GetUserById(id)
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

// 获取用户列表
func (u *UserController) GetUserList(c *gin.Context) {
	var payload GetUserListPayload
	err := c.ShouldBindQuery(&payload)
	if err != nil {
		response.FailJsonResp(c, "查询用户参数不正确")
		return
	}
	page := payload.Page
	size := payload.Size
	// conditions := map[string]interface{}{}
	result, total, err := u.UserSvc.GetUserQueryset(page, size, nil)
	if err != nil {
		response.FailJsonResp(c, "查询用户列表数据失败")
		return
	}
	response.SuccessJsonResp(c, result, map[string]interface{}{
		"total": total, "pages": utils.TotalPage(size, total),
	})
}
