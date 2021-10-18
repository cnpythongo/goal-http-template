package liveness

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"

	"github.com/cnpythongo/goal/pkg/response"
)

type ILivenessController interface {
	Ping(c *gin.Context)
}

type LivenessController struct {
}

func (l *LivenessController) Ping(c *gin.Context) {
	response.SuccessJsonResp(c, fmt.Sprintf("hello, %v", os.Getenv("APP_ROUTER_SERVICE")), nil)
}
