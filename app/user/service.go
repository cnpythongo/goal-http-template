package user

import "github.com/cnpythongo/goal/model"

type IUserService interface {
	// 根据UUID获取用户
	GetUserByUuid(uuid string) (*model.User, error)
	// 根据条件获取单一用户
	GetUserByCondition(condition interface{}) (*model.User, error)
	// 根据username获取用户
	GetUserByUsername(username string) (*model.User, error)
	// 根据email获取用户
	GetUserByEmail(email string) (*model.User, error)
}

type UserService struct {
	UserRepo IUserRepository `inject:"UserRepo"`
}

func (u *UserService) GetUserByCondition(condition interface{}) (*model.User, error) {
	result, err := u.UserRepo.GetUserByCondition(condition)
	return result, err
}

func (u *UserService) GetUserByUsername(username string) (*model.User, error) {
	result, err := u.UserRepo.GetUserByUsername(username)
	return result, err
}

func (u *UserService) GetUserByEmail(email string) (*model.User, error) {
	result, err := u.UserRepo.GetUserByEmail(email)
	return result, err
}

func (u *UserService) GetUserByUuid(uuid string) (*model.User, error) {
	result, err := u.UserRepo.GetUserByUuid(uuid)
	return result, err
}
