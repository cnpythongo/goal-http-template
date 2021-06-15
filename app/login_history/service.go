package login_history

import "github.com/cnpythongo/goal/model"

type ILoginHistoryService interface {
	GetUserLoginHistoryList(userId, page, size int) ([]*model.LoginHistory, error)
}

type LoginHistoryService struct {
	LoginHistoryRepo ILoginHistoryRepository `inject:"LoginHistoryRepo"`
}

func (l *LoginHistoryService) GetUserLoginHistoryList(userId, page, size int) ([]*model.LoginHistory, error) {
	panic("implement me")
}