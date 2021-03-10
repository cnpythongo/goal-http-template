package test

import (
	"fmt"
	"github.com/cnpythongo/goal/repo"
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestGetUserQueryset(t *testing.T) {
	users := repo.GetUserQueryset("username like ?", "%ly%")
	assert.NotEqual(t, len(*users), 0)
}

func TestGetUserObject(t *testing.T) {
	user := repo.GetUserObject(1)
	assert.NotEqual(t, user, nil)
	fmt.Println(user)
}
