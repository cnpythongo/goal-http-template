package config

import (
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"time"
)

type IAppCache interface {
	GetValue(key string) string
	SetValue(key string, value interface{}, expiration time.Duration) bool
}

type AppCache struct {
	Client *redis.Client  `inject:""`
	Logger *logrus.Logger `inject:""`
}

func (a *AppCache) GetValue(key string) string {
	val, err := a.Client.Get(key).Result()
	if err != nil {
		a.Logger.Errorln(err)
		val = ""
	}
	return val
}

func (a *AppCache) SetValue(key string, value interface{}, expiration time.Duration) bool {
	err := a.Client.Set(key, value, expiration).Err()
	if err != nil {
		a.Logger.Errorln(err)
		return false
	}
	return true
}
