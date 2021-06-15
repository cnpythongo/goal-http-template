package user_profile

import "github.com/cnpythongo/goal/model"

type IUserProfileController interface {
	// 根据用户ID获取用户资料
	GetUserProfileObjectByUserId(userId int) (*model.UserProfile, error)
}

type UserProfileController struct {
	UserProfileSvc IUserProfileService `inject:"UserProfileSvc"`
}

func (u *UserProfileController) GetUserProfileObjectByUserId(userId int) (*model.UserProfile, error) {
	panic("implement me")
}
