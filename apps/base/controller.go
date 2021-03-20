package base

import (
	"github.com/cnpythongo/goal/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

func JsonResp(c *gin.Context, code int, result interface{}, extends map[string]interface{}) {
	statusCode := http.StatusOK

	if code != common.SuccessCode {
		c.JSON(statusCode, gin.H{
			"code": code,
			"msg":  result.(string),
		})
	} else {
		jsonData := gin.H{
			"code":   code,
			"msg":    common.GetCodeMsg(code),
			"result": result,
		}
		if extends != nil {
			for key := range extends {
				jsonData[key] = extends[key]
			}
		}
		c.JSON(statusCode, jsonData)
	}
}

func Ping(c *gin.Context) {
	JsonResp(c, common.SuccessCode, map[string]string{"message": "pong"}, nil)
}
