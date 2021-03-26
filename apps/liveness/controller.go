package liveness

import (
	"github.com/gin-gonic/gin"

	"github.com/cnpythongo/goal/pkg/response"
)

func Ping(c *gin.Context) {
	response.SuccessJsonResp(c, "ok", nil)
}
