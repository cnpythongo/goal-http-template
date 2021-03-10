package test

import (
	"encoding/json"
	"github.com/go-playground/assert/v2"
	"net/http"
	"testing"
)

func TestPing(t *testing.T) {
	w := GoRequest("GET", "/ping")
	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	_ = json.Unmarshal([]byte(w.Body.String()), &response)
	result, ok := response["code"]
	result = result.(float64)
	assert.Equal(t, ok, true)
	assert.Equal(t, result, float64(0))
}
