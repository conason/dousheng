package utils

import (
	"context"
	"dousheng/config"
	redis "github.com/go-redis/redis/v8"
)

var (
	Ctx   = context.Background()
	Redis *redis.Client
)

func RedisInit() {
	Redis = redis.NewClient(&redis.Options{
		Addr: config.REDISADDR,
		DB:   config.REDISDB,
	})
}
