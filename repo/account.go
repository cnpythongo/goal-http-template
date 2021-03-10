package repo

import (
	"github.com/cnpythongo/goal/config"
	"github.com/cnpythongo/goal/model"
)

var db = config.GoalDB
var logger = config.GoalLogger

func GetUserQueryset(query interface{}, args ...interface{}) *[]model.User {
	result := model.NewUsers()
	err := db.Where(query, args...).Find(&result).Error
	if err != nil {
		logger.Errorf("dao.account.GetUserQueryset Error ==> ", err)
		return nil
	}
	return result
}

func GetUserObject(userID int) *model.User {
	result := model.NewUser()
	err := db.First(&result, userID).Error
	if err != nil {
		logger.Errorf("dao.account.GetUserObject Error ==> ", err)
		return nil
	}
	return result
}

func GetUserProfileQueryset(query interface{}, args ...interface{}) *[]model.UserProfile {
	result := model.NewUserProfiles()
	err := db.Where(query, args...).Find(&result).Error
	if err != nil {
		logger.Errorf("dao.account.GetUserProfileQueryset Error ==> ", err)
		return nil
	}
	return result
}

func GetUserProfileObject(userID int) *model.UserProfile {
	result := model.NewUserProfile()
	err := db.First(&result, userID).Error
	if err != nil {
		logger.Errorf("dao.account.GetUserProfileObject Error ==> ", err)
		return nil
	}
	return result
}

func GetLoginHistoryQueryset(query interface{}, args ...interface{}) *[]model.LoginHistory {
	result := model.NewLoginHistories()
	err := db.Where(query, args...).Find(&result).Error
	if err != nil {
		logger.Errorf("dao.account.GetLoginHistoryQueryset Error ==> ", err)
		return nil
	}
	return result
}

func GetLoginHistoryObject(id int) *model.LoginHistory {
	result := model.NewLoginHistory()
	err := db.First(&result, id).Error
	if err != nil {
		logger.Errorf("dao.account.GetLoginHistoryObject Error ==> ", err)
		return nil
	}
	return result
}
