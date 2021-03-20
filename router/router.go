package router

import (
	"github.com/cnpythongo/goal/apps/account"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"

	"github.com/cnpythongo/goal/apps/base"
)

func SetupRouters(route *gin.Engine) *gin.Engine {

	userController := account.InjectUserController()

	// middleware
	route.Use(CORSMiddleware())

	// ino project api
	userGroup := route.Group("/api/users")
	userGroup.POST("", userController.CreateUser)
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
