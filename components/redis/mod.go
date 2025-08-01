package redis

import (
	"cmc-server/resp"
	"context"
	"time"

	"github.com/beego/beego/v2/server/web"
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

type RedisService struct {
	web.Controller
}

func (r *RedisService) SetCaptcha(ctx context.Context, user string, id string, code string) (bool, error) {
	key := "captcha." + id
	userKey := "captcha." + user

	// 检测这个用户是否一分钟内发送过
	exists, err := Engine.Exists(ctx, userKey).Result()
	if err != nil {
		return false, err
	}

	// 如果已存在
	if exists > 0 {
		return false, &resp.Error{Code: 11201, Msg: "It's been less than a minute since the last request was sent"}
	}

	Engine.Set(ctx, userKey, code, 1*time.Second)
	Engine.Set(ctx, key, code, 15*time.Minute)

	return true, nil
}
func (r *RedisService) ValidateCaptcha(ctx context.Context, id string, code string) (bool, error) {
	_code, err := Engine.Get(ctx, "captcha."+id).Result()

	if err != nil {
		return false, nil
	}
	return _code == code, err
}
