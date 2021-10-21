package config

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/joho/godotenv"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io"
	"os"
	"strconv"
	"time"
)

const databaseParams = "?charset=utf8mb4&parseTime=True&loc=Local"

type Config struct {
	EnvFileName string
	Debug       bool
	Logger      *logrus.Logger
	Redis       *redis.Client
	DB          *gorm.DB
}

var GlobalConfig *Config

func init() {
	conf := &Config{}

	// GOAL_ENV_FILE env配置文件名称
	envFileName := os.Getenv("GOAL_ENV_FILE")
	if envFileName == "" {
		envFileName = ".env"
	}
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

	conf.EnvFileName = envFileName
	conf.Debug = os.Getenv("DEBUG") == "true"

	initLogger(conf)
	initDBConn(conf)
	initRedis(conf)
	GlobalConfig = conf

	if conf.Debug {
		migrateTables(conf)
	}
}

func initLogger(conf *Config) {
	logDir := os.Getenv("LOG_DIR")

	// logger setting
	conf.Logger = logrus.New()
	if conf.Debug {
		conf.Logger.SetLevel(logrus.DebugLevel)
	}
	conf.Logger.Infoln(fmt.Sprintf("Global Config Logger.Level == %d \n", conf.Logger.Level))

	logName := fmt.Sprintf("%s/%s.log", logDir, os.Getenv("GOAL_APP_SERVICE"))
	rotaLogs, _ := rotatelogs.New(logName + ".%Y%m%d")
	mw := io.MultiWriter(os.Stdout, rotaLogs)
	conf.Logger.SetOutput(mw)
	conf.Logger.SetFormatter(&logrus.JSONFormatter{})
}

func getDatabaseDsn() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
	)
}

func initDatabase() {
	dsn := fmt.Sprintf("%s%s", getDatabaseDsn(), databaseParams)
	db, err := gorm.Open(mysql.Open(dsn), nil)
	if err != nil {
		panic("Connect Database error >>> " + err.Error())
	}

	createSQL := fmt.Sprintf(
		"CREATE DATABASE IF NOT EXISTS `%s` CHARACTER SET 'utf8mb4' COLLATE 'utf8mb4_general_ci';",
		os.Getenv("DB_NAME"),
	)

	err = db.Exec(createSQL).Error
	if err != nil {
		panic("Create Database error >>> " + err.Error())
	}
}

func initDBConn(conf *Config) {
	// init database
	initDatabase()

	dsn := fmt.Sprintf("%s%s%s", getDatabaseDsn(), os.Getenv("DB_NAME"), databaseParams)
	conf.Logger.Infoln(dsn)

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

	if conf.Debug {
		conf.DB = db.Debug()
	} else {
		conf.DB = db
	}
}

func initRedis(conf *Config) {
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
	conf.Redis = redis.NewClient(options)
	if conf.Redis != nil {
		conf.Redis.Set("GOAL_SERVICE_START_AT", time.Now().Unix(), 0)
	}
}
