package user

import (
	"github.com/cnpythongo/goal-tools/utils"
	"github.com/cnpythongo/goal/config"
	"github.com/cnpythongo/goal/model"
	"github.com/cnpythongo/goal/pkg/response"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
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
	UserSvc IUserService `inject:"UserSvc"`
}

func (u *UserController) CreateUser(c *gin.Context) {
	payload := model.NewUser()
	err := c.ShouldBindJSON(payload)
	if err != nil {
		response.FailJsonResp(c, response.PayloadError, nil)
		return
	}
	eu, _ := u.UserSvc.GetUserByUsername(payload.Username)
	if eu != nil {
		response.FailJsonResp(c, response.AccountUserExistError, nil)
		return
	}
	ue, _ := u.UserSvc.GetUserByEmail(payload.Email)
	if ue != nil {
		response.FailJsonResp(c, response.AccountEmailExistsError, nil)
		return
	}
	user, err := u.UserSvc.CreateUser(payload)
	if err != nil {
		response.FailJsonResp(c, response.AccountCreateError, nil)
		return
	}
	response.SuccessJsonResp(c, user, nil)
}

func (u *UserController) GetUserById(c *gin.Context) {
	pk := c.Param("id")
	id, e := strconv.Atoi(pk)
	if e != nil {
		response.FailJsonResp(c, response.AccountUserIdError, nil)
		return
	}
	result, err := u.UserSvc.GetUserById(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			response.FailJsonResp(c, response.AccountUserNotExistError, nil)
		} else {
			response.FailJsonResp(c, response.AccountQueryUserError, nil)
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
			response.FailJsonResp(c, response.AccountUserNotExistError, nil)
		} else {
			response.FailJsonResp(c, response.AccountQueryUserError, nil)
		}
		return
	}
	response.SuccessJsonResp(c, result, nil)
}

// 获取用户列表
func (u *UserController) GetUserList(c *gin.Context) {
	var payload ReqGetUserListPayload
	err := c.ShouldBindQuery(&payload)
	if err != nil {
		config.GlobalLogger.Error(err)
		response.FailJsonResp(c, response.AccountQueryUserParamError, nil)
		return
	}
	page := payload.Page
	size := payload.Size
	// conditions := map[string]interface{}{}
	result, total, err := u.UserSvc.GetUserQueryset(page, size, nil)
	if err != nil {
		response.FailJsonResp(c, response.AccountQueryUserListError, nil)
		return
	}
	response.SuccessJsonResp(c, result, map[string]interface{}{
		"total": total, "pages": utils.TotalPage(size, total),
	})
}
