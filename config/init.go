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
)

var Debug bool
var GoalLogger *logrus.Logger
var GoalDB *gorm.DB
var GoalRedis *redis.Client

func init() {
	err := godotenv.Load()
	if err != nil {
		err := godotenv.Load("../.env")
		if err != nil {
			panic("Load .env file error >>> " + err.Error())
		}
	}

	Debug = os.Getenv("DEBUG") == "true"

	initLogger()
	initDBConn()
	initRedis()
}

func initLogger() {
	// logger setting
	GoalLogger = logrus.New()
	GoalLogger.SetLevel(logrus.DebugLevel)
	GoalLogger.Infoln("%d \n", GoalLogger.Level)
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
	fmt.Println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("database connection init error >>> " + err.Error())
	}
	GoalDB = db
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
	GoalRedis = redis.NewClient(options)
}
