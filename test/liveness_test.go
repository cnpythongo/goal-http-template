package test

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"net/http"
	"testing"

	"github.com/cnpythongo/goal/router"
	"github.com/cnpythongo/goal/test/utils"
)

func TestPing(t *testing.T) {
	r := gin.New()
	r = router.InitAPIRouters(r)
	w := utils.DoRequest(r, "GET", "/api/ping", nil)
	assert.Equal(t, http.StatusOK, w.Code)
	response := utils.ParseResponseToJSON(w)
	result, ok := response["code"]
	result = result.(float64)
	assert.Equal(t, ok, true)
	assert.Equal(t, result, float64(1000))
}
