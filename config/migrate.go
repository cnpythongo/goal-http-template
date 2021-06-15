package config

import (
	"github.com/cnpythongo/goal/model"
)

func migrateTables() {
	if !Debug { // 仅在开发模式执行migrate操作
		return
	}
	GlobalLogger.Infoln("migrate tables start .....")
	err := GlobalDB.AutoMigrate(model.NewUser())
	if err != nil {
		panic(err)
	}
	err = GlobalDB.AutoMigrate(model.NewUserProfile())
	if err != nil {
		panic(err)
	}
	err = GlobalDB.AutoMigrate(model.NewLoginHistory())
	if err != nil {
		panic(err)
	}
	GlobalLogger.Infoln("migrate tables success .....")
}
