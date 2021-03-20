package account

import (
	"github.com/cnpythongo/goal/apps/account/controller"
	"github.com/cnpythongo/goal/apps/account/repository"
	"github.com/cnpythongo/goal/apps/account/service"
	"github.com/facebookgo/inject"
)

func InjectUserController() controller.UserController {
	var ctl controller.UserController
	var injector inject.Graph
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

func InjectUserProfileController() controller.UserProfileController {
	var ctl controller.UserProfileController
	var injector inject.Graph
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

func InjectLoginHistoryController() controller.LoginHistoryController {
	var ctl controller.LoginHistoryController
	var injector inject.Graph
	err := injector.Provide(
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
