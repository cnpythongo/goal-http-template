package admin

import (
	"github.com/facebookgo/inject"

	"github.com/cnpythongo/goal/config"

	"github.com/cnpythongo/goal/admin/login_history"
	"github.com/cnpythongo/goal/admin/user"
	"github.com/cnpythongo/goal/admin/user_profile"
)

func InjectUserController(injector inject.Graph) user.UserController {
	var ctl user.UserController
	err := injector.Provide(
		&inject.Object{Value: &user.UserRepository{}, Name: "UserRepo"},
		&inject.Object{Value: &user.UserService{}, Name: "UserSvc"},
		&inject.Object{Value: &ctl},
	)
	if err != nil {
		panic("inject fatal: " + err.Error())
	}
	if err := injector.Populate(); err != nil {
		panic("inject fatal: " + err.Error())
	}
	return ctl
}

func InjectUserProfileController(injector inject.Graph) user_profile.UserProfileController {
	var ctl user_profile.UserProfileController
	err := injector.Provide(
		&inject.Object{Value: &user_profile.UserProfileRepository{}, Name: "UserProfileRepo"},
		&inject.Object{Value: &user_profile.UserProfileService{}, Name: "UserProfileSvc"},
		&inject.Object{Value: &ctl},
	)
	if err != nil {
		panic("inject fatal: " + err.Error())
	}
	if err := injector.Populate(); err != nil {
		panic("inject fatal: " + err.Error())
	}
	return ctl
}

func InjectLoginHistoryController(injector inject.Graph) login_history.LoginHistoryController {
	var ctl login_history.LoginHistoryController
	err := injector.Provide(
		&inject.Object{Value: config.GlobalDB},
		&inject.Object{Value: config.GlobalLogger},
		&inject.Object{Value: &login_history.LoginHistoryRepository{}, Name: "LoginHistoryRepo"},
		&inject.Object{Value: &login_history.LoginHistoryService{}, Name: "LoginHistorySvc"},
		&inject.Object{Value: &ctl},
	)
	if err != nil {
		panic("inject fatal: " + err.Error())
	}
	if err := injector.Populate(); err != nil {
		panic("inject fatal: " + err.Error())
	}
	return ctl
}
