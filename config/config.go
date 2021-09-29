package config

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"strconv"
	"time"
)

var Debug bool
var GlobalLogger *logrus.Logger
var GlobalDB *gorm.DB
var GlobalRedis *redis.Client
var GlobalEnvName string

func init() {
	// ENV_NAME env配置文件名称
	GlobalEnvName = os.Getenv("ENV_NAME")
	if GlobalEnvName == "" {
		GlobalEnvName = ".env.local"
	}

	envFileName := GlobalEnvName
	// 跑测试用例时找不到配置文件,最多向上找20层目录
	for i := 0; i < 20; i++ {
		if _, err := os.Stat(envFileName); err == nil {
			break
		} else {
			envFileName = "../" + envFileName
		}
	}

	err := godotenv.Load(envFileName)
	if err != nil {
		panic("Load .env file error >>> " + err.Error())
	}

	Debug = os.Getenv("DEBUG") == "true"

	initLogger()
	initDBConn()
	initRedis()
	if Debug {
		migrateTables()
	}
}

func initLogger() {
	// logger setting
	GlobalLogger = logrus.New()
	GlobalLogger.SetLevel(logrus.DebugLevel)
	GlobalLogger.Infoln(fmt.Sprintf("GlobalLogger.Level == %d \n", GlobalLogger.Level))
}

func initDBConn() {
	// db setting
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	GlobalLogger.Infoln(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("database connection init error >>> " + err.Error())
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic("database sql db instance init error >>> " + err.Error())
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Minute * 30)

	if Debug == true {
		GlobalDB = db.Debug()
	} else {
		GlobalDB = db
	}
}

func initRedis() {
	if os.Getenv("ENABLE_REDIS") == "false" {
		return
	}
	host := fmt.Sprintf("%s:%s",
		os.Getenv("REDIS_HOST"),
		os.Getenv("REDIS_PORT"),
	)
	db, _ := strconv.Atoi(os.Getenv("REDIS_DB"))
	pwd := os.Getenv("REDIS_PASSWORD")
	options := &redis.Options{
		Addr: host,
		DB:   db,
	}
	if pwd != "" {
		options.Password = pwd
	}
	GlobalRedis = redis.NewClient(options)
}
