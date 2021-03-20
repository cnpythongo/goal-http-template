package repository

import (
	"github.com/cnpythongo/goal/apps/account/model"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type IUserRepository interface {
	// 创建用户
	CreateUser(user *model.User) (*model.User, error)
	// 根据ID获取用户
	GetUserById(id int) (*model.User, error)
	// 根据UUID获取用户
	GetUserByUuid(uuid string) (*model.User, error)
	// 获取用户查询集
	GetUserQueryset(page, size int, conditions interface{}) ([]*model.User, error)
}

type UserRepository struct {
	DB     *gorm.DB       `inject:""`
	Logger *logrus.Logger `inject:""`
}

func (u *UserRepository) CreateUser(user *model.User) (*model.User, error) {
	panic("implement me")
}

func (u *UserRepository) GetUserByUuid(uuid string) (*model.User, error) {
	panic("implement me")
}

func (u *UserRepository) GetUserQueryset(page, size int, conditions interface{}) ([]*model.User, error) {
	result := model.NewUsers()
	offset := (page - 1) * size
	err := u.DB.Debug().Where(conditions).Limit(size).Offset(offset).Find(&result).Error
	if err != nil {
		u.Logger.Errorf("apps.account.UserRepository.GetUserQueryset Error ==> ", err)
		return nil, err
	}
	return result, nil
}

func (u *UserRepository) GetUserById(userID int) (*model.User, error) {
	result := model.NewUser()
	err := u.DB.Debug().First(&result, userID).Error
	if err != nil {
		u.Logger.Errorf("apps.account.UserRepository.GetUserById Error ==> ", err)
		return nil, err
	}
	return result, nil
}
