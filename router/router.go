package router

import (
	"github.com/facebookgo/inject"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"

	"github.com/cnpythongo/goal/config"

	"github.com/cnpythongo/goal/apps/account"
	"github.com/cnpythongo/goal/apps/liveness"
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

func SetupRouters(route *gin.Engine) *gin.Engine {
	var injector inject.Graph
	err := injector.Provide(
		&inject.Object{Value: config.GlobalDB},
		&inject.Object{Value: config.GlobalLogger},
	)
	if err != nil {
		panic("inject fatal: " + err.Error())
	}

	userController := account.InjectUserController(injector)

	// middleware
	route.Use(CORSMiddleware())

	// ino project api
	userGroup := route.Group("/api/users")
	userGroup.POST("", userController.CreateUser)
	userGroup.GET("/:uid", userController.GetUserByUuid)

	// common test api
	apiGroup := route.Group("/api")
	apiGroup.GET("/ping", liveness.Ping)
	return route
}
