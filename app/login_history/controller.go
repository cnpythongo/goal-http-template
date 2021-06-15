package login_history

import (
	"github.com/gin-gonic/gin"
)

type ILoginHistoryController interface {
	GetUserLoginHistory(c *gin.Context)
}

type LoginHistoryController struct {
	LoginHistorySvc ILoginHistoryService `inject:"LoginHistorySvc"`
}

func (l *LoginHistoryController) GetUserLoginHistory(c *gin.Context) {
	panic("implement me")
}
