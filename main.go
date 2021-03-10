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
	route := gin.New()
	route = router.SetupRouters(route)
	addr := fmt.Sprintf("%s:%s", host, port)
	route.Run(addr)
}
