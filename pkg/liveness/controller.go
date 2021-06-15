package liveness

import (
	"github.com/gin-gonic/gin"

	"github.com/cnpythongo/goal/pkg/response"
)

type ILivenessController interface {
	Ping(c *gin.Context)
}

type LivenessController struct {
}

func (l *LivenessController) Ping(c *gin.Context) {
	response.SuccessJsonResp(c, "ok", nil)
}
