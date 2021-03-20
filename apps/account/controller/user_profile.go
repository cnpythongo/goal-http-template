package controller

import (
	"github.com/cnpythongo/goal/apps/account/model"
	"github.com/cnpythongo/goal/apps/account/service"
)

type IUserProfileController interface {
	// 根据用户ID获取用户资料
	GetUserProfileObjectByUserId(userId int) (*model.UserProfile, error)
}

type UserProfileController struct {
	UserProfileSvc service.IUserProfileService `inject:"UserProfileSvc"`
}

func (u *UserProfileController) GetUserProfileObjectByUserId(userId int) (*model.UserProfile, error) {
	panic("implement me")
}
