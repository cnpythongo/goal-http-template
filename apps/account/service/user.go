package service

import (
	"github.com/cnpythongo/goal/apps/account/model"
	"github.com/cnpythongo/goal/apps/account/repository"
)

type IUserService interface {
	// 创建用户
	CreateUser(user *model.User) (*model.User, error)
	// 根据ID获取用户
	GetUserById(id int) (*model.User, error)
	// 根据UUID获取用户
	GetUserByUuid(uuid string) (*model.User, error)
	// 获取用户查询集
	GetUserQueryset(page, size int, conditions interface{}) ([]*model.User, error)
}

type UserService struct {
	UserRepo repository.IUserRepository `inject:"UserRepo"`
}

func (u *UserService) CreateUser(user *model.User) (*model.User, error) {
	result, err := u.UserRepo.CreateUser(user)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (u *UserService) GetUserById(id int) (*model.User, error) {
	panic("implement me")
}

func (u *UserService) GetUserByUuid(uuid string) (*model.User, error) {
	result, err := u.UserRepo.GetUserByUuid(uuid)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (u *UserService) GetUserQueryset(page, size int, conditions interface{}) ([]*model.User, error) {
	panic("implement me")
}
