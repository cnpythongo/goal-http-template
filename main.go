package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"

	"github.com/cnpythongo/goal/config"
	"github.com/cnpythongo/goal/router"
)

func main() {
	if !config.GlobalConfig.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
	svc := os.Getenv("GOAL_APP_SERVICE")
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	address := fmt.Sprintf("%s:%s", host, port)
	config.GlobalConfig.Logger.Info(fmt.Sprintf("Server: %s", address))

	route := gin.New()
	if svc == "admin" {
		router.InitAdminRouters(route)
	} else {
		router.InitAPIRouters(route)
	}
	server := router.GetDefaultHttpServer(address, route)
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
