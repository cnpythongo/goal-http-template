package account

import (
	account2 "github.com/cnpythongo/goal/model/account"
	"github.com/cnpythongo/goal/repository/account"
)

type IUserService interface {
	// 创建用户
	CreateUser(user *account2.User) (*account2.User, error)
	// 根据ID获取用户
	GetUserById(id int) (*account2.User, error)
	// 根据UUID获取用户
	GetUserByUuid(uuid string) (*account2.User, error)
	// 获取用户查询集
	GetUserQueryset(page, size int, conditions interface{}) ([]*account2.User, int, error)
	// 根据条件获取单一用户
	GetUserByCondition(condition interface{}) (*account2.User, error)
	// 根据username获取用户
	GetUserByUsername(username string) (*account2.User, error)
	// 根据email获取用户
	GetUserByEmail(email string) (*account2.User, error)
}

type UserService struct {
	UserRepo account.IUserRepository `inject:"UserRepo"`
}

func (u *UserService) GetUserByCondition(condition interface{}) (*account2.User, error) {
	result, err := u.UserRepo.GetUserByCondition(condition)
	return result, err
}

func (u *UserService) GetUserByUsername(username string) (*account2.User, error) {
	result, err := u.UserRepo.GetUserByUsername(username)
	return result, err
}

func (u *UserService) GetUserByEmail(email string) (*account2.User, error) {
	result, err := u.UserRepo.GetUserByEmail(email)
	return result, err
}

func (u *UserService) CreateUser(user *account2.User) (*account2.User, error) {
	result, err := u.UserRepo.CreateUser(user)
	return result, err
}

func (u *UserService) GetUserById(id int) (*account2.User, error) {
	result, err := u.UserRepo.GetUserById(id)
	return result, err
}

func (u *UserService) GetUserByUuid(uuid string) (*account2.User, error) {
	result, err := u.UserRepo.GetUserByUuid(uuid)
	return result, err
}

func (u *UserService) GetUserQueryset(page, size int, conditions interface{}) ([]*account2.User, int, error) {
	result, total, err := u.UserRepo.GetUserQueryset(page, size, conditions)
	return result, total, err
}
