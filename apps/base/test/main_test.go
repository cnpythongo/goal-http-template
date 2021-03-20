package test

import (
	"bytes"
	"encoding/json"
	"github.com/cnpythongo/goal/router"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"net/http"
	"net/http/httptest"
	"testing"
)

func DoRequest(r http.Handler, method, path string, data interface{}) *httptest.ResponseRecorder {
	var reqBody *bytes.Buffer
	if data != nil {
		body, _ := json.Marshal(data)
		reqBody = bytes.NewBuffer(body)
	} else {
		body, _ := json.Marshal(make(map[string]string))
		reqBody = bytes.NewBuffer(body)
	}
	req, _ := http.NewRequest(method, path, reqBody)
	w := httptest.NewRecorder()
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	return w
}

func ParseResponseToJSON(w *httptest.ResponseRecorder) map[string]interface{} {
	var response map[string]interface{}
	_ = json.Unmarshal([]byte(w.Body.String()), &response)
	return response
}

func TestPing(t *testing.T) {
	r := gin.New()
	r = router.SetupRouters(r)
	w := DoRequest(r, "GET", "/api/ping", nil)
	assert.Equal(t, http.StatusOK, w.Code)
	response := ParseResponseToJSON(w)
	result, ok := response["code"]
	result = result.(float64)
	assert.Equal(t, ok, true)
	assert.Equal(t, result, float64(1000))
}


//func TestCreateProject(t *testing.T) {
//	payload := &ino.ProjectPayload{
//		AuthSignature: "0x4b9eE53b17bCa029FE2Bb2A9aE13321ba",
//		Email:         "aaa@qq.com",
//		Logo:          "https://aaaa.jpg",
//		Name:          "币升项目",
//		Desc:          "币升项目币升项目币升项目币升项目的简介",
//		ContractType:  "ETH",
//		ContractAddr:  "xxxsdfsdfsdfsdf",
//		WebSite:       "https://gogog.com",
//		WhitePaper:    "https://sdfsdfsd.com",
//		Twitter:       "sdfsdf@fads.com",
//		Telegram:      "xxcsdfsdf",
//		Discord:       "dfsdfsdf",
//	}
//	r := gin.New()
//	r = router.SetupRouters(r)
//	w := DoRequest(r, "POST", "/api/ino/projects", payload)
//	assert.Equal(t, http.StatusOK, w.Code)
//
//	response := ParseResponseToJSON(w)
//	result, ok := response["code"]
//	result = result.(float64)
//	assert.Equal(t, ok, true)
//	assert.Equal(t, result, float64(1000))
//}
