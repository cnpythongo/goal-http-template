package admin

import (
	"github.com/facebookgo/inject"

	"github.com/cnpythongo/goal/config"

	"github.com/cnpythongo/goal/repository"
	"github.com/cnpythongo/goal/service"
)

func InjectUserController(injector inject.Graph) UserController {
	var ctl UserController
	err := injector.Provide(
		&inject.Object{Value: &repository.UserRepository{}, Name: "UserRepo"},
		&inject.Object{Value: &service.UserService{}, Name: "UserSvc"},
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

func InjectUserProfileController(injector inject.Graph) UserProfileController {
	var ctl UserProfileController
	err := injector.Provide(
		&inject.Object{Value: &repository.UserProfileRepository{}, Name: "UserProfileRepo"},
		&inject.Object{Value: &service.UserProfileService{}, Name: "UserProfileSvc"},
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

func InjectLoginHistoryController(injector inject.Graph) LoginHistoryController {
	var ctl LoginHistoryController
	err := injector.Provide(
		&inject.Object{Value: config.GlobalDB},
		&inject.Object{Value: config.GlobalLogger},
		&inject.Object{Value: &repository.LoginHistoryRepository{}, Name: "LoginHistoryRepo"},
		&inject.Object{Value: &service.LoginHistoryService{}, Name: "LoginHistorySvc"},
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
