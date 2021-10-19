package config

import (
	"github.com/cnpythongo/goal/model"
)

func migrateTables(conf *Config) {
	if !GlobalConfig.Debug { // 仅在开发模式执行migrate操作
		return
	}
	GlobalConfig.Logger.Infoln("migrate tables start .....")
	err := GlobalConfig.DB.AutoMigrate(model.NewUser())
	if err != nil {
		panic(err)
	}
	err = GlobalConfig.DB.AutoMigrate(model.NewUserProfile())
	if err != nil {
		panic(err)
	}
	err = GlobalConfig.DB.AutoMigrate(model.NewLoginHistory())
	if err != nil {
		panic(err)
	}
	GlobalConfig.Logger.Infoln("migrate tables success .....")
}
