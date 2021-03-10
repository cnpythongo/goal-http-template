package handler

import (
	"github.com/cnpythongo/goal/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

func JsonResp(c *gin.Context, code int, result interface{}) {
	statusCode := http.StatusOK

	if code != common.SuccessCode {
		statusCode = http.StatusBadRequest
	}

	c.JSON(statusCode, gin.H{
		"code":   code,
		"msg":    common.GetCodeMsg(code),
		"result": result,
	})
}

func Ping(c *gin.Context) {
	JsonResp(c, common.SuccessCode, map[string]string{"message": "pong"})
}
