package redis

import (
	"github.com/mouuii/otccommon/config"

	"github.com/go-redis/redis"
)

var redisClient *redisClient

func InitRedis(cfg config.RedisCfg) error {
	option := &redis.Options{
		Addr:     cfg.ConnectionURI,
		Password: cfg.Pwd,
	}
	redisClient = redis.NewClient(option)

	if _, err := redisClient.Ping().Result(); err != nil {
		return err
	}
	return nil
}
