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

func init() {
	// APP_RUN_ENV 系统环境变量，dev：开发环境, test: 测试环境，pro: 生产环境
	runEnv := os.Getenv("APP_RUN_ENV")
	if runEnv == "" {
		runEnv = "dev"
	}
	envFileName := fmt.Sprintf(".env.%s", runEnv)
	err := godotenv.Load(envFileName)
	if err != nil {
		envFileName = fmt.Sprintf("../../../.env.%s", runEnv)
		err := godotenv.Load(envFileName)
		if err != nil {
			panic("Load .env file error >>> " + err.Error())
		}
	}

	Debug = os.Getenv("DEBUG") == "true"

	initLogger()
	initDBConn()
	initRedis()
	migrateTables()
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

	GlobalDB = db
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
