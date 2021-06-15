package router

import (
	"github.com/cnpythongo/goal/admin"
	"github.com/facebookgo/inject"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"

	"github.com/cnpythongo/goal/config"

	"github.com/cnpythongo/goal/app"
	"github.com/cnpythongo/goal/pkg/liveness"
)

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

func InitAPIRouters(route *gin.Engine) *gin.Engine {
	var injector inject.Graph
	err := injector.Provide(
		&inject.Object{Value: config.GlobalDB},
		&inject.Object{Value: config.GlobalLogger},
	)
	if err != nil {
		panic("inject fatal: " + err.Error())
	}

	userController := app.InjectUserController(injector)
	liveController := liveness.InjectLivenessController(injector)

	// middleware
	route.Use(CORSMiddleware())
	// common test api
	apiGroup := route.Group("/api")
	apiGroup.GET("/ping", liveController.Ping)
	// api
	userGroup := route.Group("/api/account")
	userGroup.GET("/me", userController.GetUserByUuid)
	userGroup.GET("/users/:uid", userController.GetUserByUuid)
	return route
}

func InitAdminRouters(route *gin.Engine) *gin.Engine {
	var injector inject.Graph
	err := injector.Provide(
		&inject.Object{Value: config.GlobalDB},
		&inject.Object{Value: config.GlobalLogger},
	)
	if err != nil {
		panic("inject fatal: " + err.Error())
	}

	userController := admin.InjectUserController(injector)
	liveController := liveness.InjectLivenessController(injector)

	// middleware
	route.Use(CORSMiddleware())

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
