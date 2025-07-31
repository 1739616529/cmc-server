package redis

import (
	redis "github.com/redis/go-redis/v9"
)

var (
	Engine *redis.Client
)

func Init() {
	Engine = redis.NewClient(&redis.Options{
		Addr:     "redis-17095.crce178.ap-east-1-1.ec2.redns.redis-cloud.com:17095", // Redis 地址
		Password: "TQGsNGsGbyOprYPl09wYVQZlYqvs7u5l",                                // 没有密码则留空
		DB:       0,                                                                 // 默认 DB
	})
}
