package user

import (
	"github.com/cnpythongo/goal/model"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type IUserRepository interface {
	// 根据UUID获取用户
	GetUserByUuid(uuid string) (*model.User, error)
	// 根据条件获取单一用户
	GetUserByCondition(condition interface{}) (*model.User, error)
	// 根据username获取用户
	GetUserByUsername(username string) (*model.User, error)
	// 根据email获取用户
	GetUserByEmail(email string) (*model.User, error)
}

type UserRepository struct {
	DB     *gorm.DB       `inject:""`
	Logger *logrus.Logger `inject:""`
}

func (u *UserRepository) GetUserByCondition(condition interface{}) (*model.User, error) {
	result := model.NewUser()
	err := u.DB.Where(condition).First(result).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			u.Logger.Errorf("apps.account.UserRepository.GetUserByCondition Error ==> %v", err)
			u.Logger.Infof("condition ==> %v", condition)
		}
		return nil, err
	}
	return result, nil
}

func (u *UserRepository) GetUserByUsername(username string) (*model.User, error) {
	condition := map[string]interface{}{"username": username}
	result, err := u.GetUserByCondition(condition)
	return result, err
}

func (u *UserRepository) GetUserByEmail(email string) (*model.User, error) {
	condition := map[string]interface{}{"email": email}
	result, err := u.GetUserByCondition(condition)
	return result, err
}

func (u *UserRepository) GetUserByUuid(uuid string) (*model.User, error) {
	condition := map[string]interface{}{"uuid": uuid}
	result, err := u.GetUserByCondition(condition)
	return result, err
}
