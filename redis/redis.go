package redis

import (
	"github.com/mouuii/otccommon/config"

	"github.com/go-redis/redis"
)

var RedisClient *redis.Client

func InitRedis(cfg config.RedisCfg) error {
	option := &redis.Options{
		Addr:     cfg.ConnectionURI,
		Password: cfg.Pwd,
	}
	RedisClient = redis.NewClient(option)

	if _, err := RedisClient.Ping().Result(); err != nil {
		return err
	}
	return nil
}
