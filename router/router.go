package router

import (
	accountRepo "github.com/cnpythongo/goal/apps/account/repository"
	"github.com/facebookgo/inject"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"

	"github.com/cnpythongo/goal/apps/base"
	"github.com/cnpythongo/goal/config"

	accountController "github.com/cnpythongo/goal/apps/account/controller"
	accountSvc "github.com/cnpythongo/goal/apps/account/service"
)

func SetupRouters(route *gin.Engine) *gin.Engine {

	var userController accountController.UserController
	var userProfileController accountController.UserProfileController
	var loginHistoryController accountController.LoginHistoryController

	// Injection，
	// Tips: 必须给注入的repository和service设置Name别名, 与struct里inject标签保持一致，否则会注入失败
	var injector inject.Graph
	err := injector.Provide(
		&inject.Object{Value: config.GlobalDB},
		&inject.Object{Value: config.GlobalLogger},

		&inject.Object{Value: &accountRepo.UserRepository{}, Name: "UserRepo"},
		&inject.Object{Value: &accountSvc.UserService{}, Name: "UserSvc"},
		&inject.Object{Value: &userController},

		&inject.Object{Value: &accountRepo.UserProfileRepository{}, Name: "UserProfileRepo"},
		&inject.Object{Value: &accountSvc.UserProfileService{}, Name: "UserProfileSvc"},
		&inject.Object{Value: &userProfileController},

		&inject.Object{Value: &accountRepo.LoginHistoryRepository{}, Name: "LoginHistoryRepo"},
		&inject.Object{Value: &accountSvc.LoginHistoryService{}, Name: "LoginHistorySvc"},
		&inject.Object{Value: &loginHistoryController},
	)
	if err != nil {
		panic("inject fatal: " + err.Error())
	}
	if err := injector.Populate(); err != nil {
		panic("inject fatal: " + err.Error())
	}

	// middleware
	route.Use(CORSMiddleware())

	// ino project api
	userGroup := route.Group("/api/users")
	userGroup.GET("/:uuid", userController.GetUserByUuid)

	// common test api
	apiGroup := route.Group("/api")
	apiGroup.GET("/ping", base.Ping)
	return route
}

func GetDefaultHttpServer(addr string, route *gin.Engine) *http.Server {
	return &http.Server{
		Addr:              addr,
		Handler:           route,
		ReadTimeout:       100 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      60 * time.Second,
		MaxHeaderBytes:    1 << 20,
	}
}
