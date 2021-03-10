package test

import (
	"github.com/cnpythongo/goal/router"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
)

func GoRequest(method, path string) *httptest.ResponseRecorder {
	route := gin.New()
	route = router.SetupRouters(route)
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	route.ServeHTTP(w, req)
	return w
}
