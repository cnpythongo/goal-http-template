package test

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"net/http"
	"testing"

	_ "github.com/cnpythongo/goal/config"

	"github.com/cnpythongo/goal/apps/account/model"
	"github.com/cnpythongo/goal/pkg/basic"
	"github.com/cnpythongo/goal/router"
)

func getRouter() *gin.Engine {
	r := gin.New()
	r = router.SetupRouters(r)
	return r
}

func TestCreateUser(t *testing.T) {
	payload := model.User{
		Username: "lyh2",
		Password: "123123",
		Email:    "aaabbb@qq.com",
		Avatar:   "http://www.qq.com/aaa.jpg",
	}
	r := getRouter()
	w := basic.DoRequest(r, "POST", "/api/users", payload)
	assert.Equal(t, http.StatusOK, w.Code)

	response := basic.ParseResponseToJSON(w)
	result, ok := response["code"]
	fmt.Printf("%v", response)
	result = result.(float64)
	assert.Equal(t, ok, true)
	assert.Equal(t, result, float64(1000))
}

func TestGetUserByUuid(t *testing.T) {
	r := getRouter()
	uid := "3610b2e5ab0a43c6b909eece0cb1c167"
	w := basic.DoRequest(r, "GET", fmt.Sprintf("/api/users/%s", uid), nil)
	assert.Equal(t, http.StatusOK, w.Code)
	response := basic.ParseResponseToJSON(w)
	fmt.Printf("%v", response)
	result, ok := response["code"]
	result = result.(float64)
	assert.Equal(t, ok, true)
	assert.Equal(t, result, float64(1000))
}
