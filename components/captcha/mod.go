package captcha

import (
	"cmc-server/components/redis"
	"cmc-server/resp"
	"context"
	"time"

	"github.com/beego/beego/v2/server/web"
)

var (
	ThrottlingTime = 1 * time.Second  // 防抖 再次发送验证码间隔时间
	ExpiredTime    = 15 * time.Minute // 验证码过期时间
)

type CaptchaService struct {
	web.Controller
}

func (r *CaptchaService) SetCaptcha(ctx context.Context, user string, id string, code string) (bool, error) {
	key := "captcha." + id
	userKey := "captcha." + user

	// 检测这个用户是否一分钟内发送过
	exists, err := redis.Engine.Exists(ctx, userKey).Result()
	if err != nil {
		return false, err
	}

	// 如果已存在
	if exists > 0 {
		return false, resp.NewError(resp.StatusCaptchaExpiration)
	}

	redis.Engine.Set(ctx, userKey, code, ThrottlingTime)
	redis.Engine.Set(ctx, key, code, ExpiredTime)

	return true, nil
}
func (r *CaptchaService) ValidateCaptcha(ctx context.Context, id string, code string) (bool, error) {
	_code, err := redis.Engine.Get(ctx, "captcha."+id).Result()

	if err != nil {
		println("111, ", err.Error())
		return false, resp.NewError(resp.StatusCaptchaExpiration)
	}
	return _code == code, nil
}
