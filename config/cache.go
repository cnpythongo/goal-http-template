package config

import (
	"time"
)

func GetValue(key string) string {
	val, err := GlobalRedis.Get(key).Result()
	if err != nil {
		GlobalLogger.Errorln(err)
		val = ""
	}
	return val
}

func SetValue(key string, value interface{}, expiration time.Duration) bool {
	err := GlobalRedis.Set(key, value, expiration).Err()
	if err != nil {
		GlobalLogger.Errorln(err)
		return false
	}
	return true
}
