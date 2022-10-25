package injectors

import (
	"github.com/cnpythongo/goal/controller/admin/account"
	"github.com/cnpythongo/goal/model"
	"github.com/cnpythongo/goal/pkg/common/log"
	account2 "github.com/cnpythongo/goal/repository/account"
	account3 "github.com/cnpythongo/goal/service/account"
	"github.com/facebookgo/inject"
)

func InjectUserController(injector inject.Graph) account.UserController {
	var ctl account.UserController
	err := injector.Provide(
		&inject.Object{Value: &account2.UserRepository{}, Name: "UserRepo"},
		&inject.Object{Value: &account3.UserService{}, Name: "UserSvc"},
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

func InjectUserProfileController(injector inject.Graph) account.UserProfileController {
	var ctl account.UserProfileController
	err := injector.Provide(
		&inject.Object{Value: &account2.UserProfileRepository{}, Name: "UserProfileRepo"},
		&inject.Object{Value: &account3.UserProfileService{}, Name: "UserProfileSvc"},
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

func InjectLoginHistoryController(injector inject.Graph) account.LoginHistoryController {
	var ctl account.LoginHistoryController
	err := injector.Provide(
		&inject.Object{Value: model.GetDB()},
		&inject.Object{Value: log.GetLogger()},
		&inject.Object{Value: &account2.LoginHistoryRepository{}, Name: "LoginHistoryRepo"},
		&inject.Object{Value: &account3.LoginHistoryService{}, Name: "LoginHistorySvc"},
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
