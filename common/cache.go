package common

import (
	"github.com/cnpythongo/goal/config"
	"time"
)

var redis = config.GoalRedis
var logger = config.GoalLogger

func GetValue(key string) string {
	val, err := redis.Get(key).Result()
	if err != nil {
		logger.Errorln(err)
		val = ""
	}
	return val
}

func SetValue(key string, value interface{}, expiration time.Duration) bool {
	err := redis.Set(key, value, expiration).Err()
	if err != nil {
		logger.Errorln(err)
		return false
	}
	return true
}
