package base

import (
	"github.com/cnpythongo/goal/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

func JsonResp(c *gin.Context, code int, result interface{}, extends interface{}) {
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
			ex := extends.(map[string]interface{})
			for key := range ex {
				jsonData[key] = ex[key]
			}
		}
		c.JSON(statusCode, jsonData)
	}
}

func SuccessJsonResp(c *gin.Context, result interface{}, extends map[string]interface{}) {
	JsonResp(c, common.SuccessCode, result, extends)
}

func FailJsonResp(c *gin.Context, message string) {
	JsonResp(c, common.FailCode, message, nil)
}

func Ping(c *gin.Context) {
	SuccessJsonResp(c, "pong", map[string]interface{}{"go": "good"})
}
