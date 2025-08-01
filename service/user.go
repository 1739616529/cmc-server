package service

import "cmc-server/components/captcha"

type UserService struct {
	captchaService captcha.CaptchaService
}
