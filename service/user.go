package service

import (
	"cmc-server/components/redis"
)

type UserService struct {
	redisService redis.RedisService
}
