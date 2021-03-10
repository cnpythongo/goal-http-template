package router

import (
	"github.com/cnpythongo/goal/handler"
	"github.com/gin-gonic/gin"
)

func SetupRouters(route *gin.Engine) *gin.Engine {
	route.GET("/ping", handler.Ping)
	return route
}
