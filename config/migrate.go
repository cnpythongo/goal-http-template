package config

import "github.com/cnpythongo/goal/apps/account/model"

func migrateTables() {
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
