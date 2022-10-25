package router

import (
	"github.com/cnpythongo/goal/controller/admin/injectors"
	"github.com/facebookgo/inject"
	"github.com/gin-gonic/gin"

	"github.com/cnpythongo/goal/controller/liveness"
	"github.com/cnpythongo/goal/model"
	"github.com/cnpythongo/goal/pkg/common/config"
	"github.com/cnpythongo/goal/pkg/common/log"
)

func InitAdminRouters(cfg *config.Configuration) *gin.Engine {
	route := initDefaultRouter(cfg)

	var injector inject.Graph
	err := injector.Provide(
		&inject.Object{Value: model.GetDB()},
		&inject.Object{Value: log.GetLogger()},
	)
	if err != nil {
		panic("inject fatal: " + err.Error())
	}

	userController := injectors.InjectUserController(injector)
	liveController := liveness.InjectLivenessController(injector)

	// common test api
	apiGroup := route.Group("/api")
	apiGroup.GET("/ping", liveController.Ping)

	// admin api
	adminGroup := route.Group("/api/account")
	adminGroup.GET("/users", userController.GetUserList)
	adminGroup.GET("/users/:uid", userController.GetUserByUuid)
	adminGroup.POST("/users", userController.CreateUser)
	return route
}
