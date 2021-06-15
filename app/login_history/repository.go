package login_history

import (
	"github.com/cnpythongo/goal/model"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ILoginHistoryRepository interface {
	GetUserLoginHistoryList(userId, page, size int) ([]*model.LoginHistory, error)
}

type LoginHistoryRepository struct {
	DB     *gorm.DB       `inject:""`
	Logger *logrus.Logger `inject:""`
}

func (l *LoginHistoryRepository) GetUserLoginHistoryList(userId, page, size int) ([]*model.LoginHistory, error) {
	panic("implement me")
}
