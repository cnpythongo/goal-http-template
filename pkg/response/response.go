package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func JsonResp(c *gin.Context, code int, result interface{}, extends interface{}) {
	statusCode := http.StatusOK

	if code != SuccessCode {
		c.JSON(statusCode, gin.H{
			"code": code,
			"msg":  result.(string),
		})
	} else {
		jsonData := gin.H{
			"code":   code,
			"msg":    GetCodeMsg(code),
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
	JsonResp(c, SuccessCode, result, extends)
}

func FailJsonResp(c *gin.Context, message string) {
	JsonResp(c, FailCode, message, nil)
}
