package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"

	"github.com/cnpythongo/goal/config"
	"github.com/cnpythongo/goal/router"
)

func main() {
	if !config.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	address := fmt.Sprintf("%s:%s", host, port)
	config.GlobalLogger.Info(fmt.Sprintf("Server: %s", address))

	route := gin.New()
	router.InitAdminRouters(route)
	server := router.GetDefaultHttpServer(address, route)
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
