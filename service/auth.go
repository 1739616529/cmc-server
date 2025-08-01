package service

import (
	"cmc-server/components/orm"
	"cmc-server/components/redis"
	"cmc-server/dto"
	"cmc-server/models"
	"cmc-server/resp"
	"cmc-server/util"
	"context"
)

type AuthService struct {
	redisService redis.RedisService
}

func (u *AuthService) Login(ctx context.Context, dto *dto.UserLogin) (*models.User, error) {

	var user models.User

	has, err := orm.Engine.Where("email = ? OR passwd = ?", dto.Account, dto.Account).Get(&user)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, resp.Error{Code: 100104, Msg: "user Not Found"}
	}

	success, err := u.redisService.ValidateCaptcha(ctx, dto.VerifyId, dto.Code)
	if err != nil {
		return nil, err
	}

	if success == false {
		return nil, resp.Error{Code: 110202, Msg: "CAPTCHA error"}
	}

	return &user, nil // 没有匹配的用户
}

func (u *AuthService) Register(ctx context.Context, dto *dto.UserRegister) (*models.User, error) {

	var user models.User
	has, err := orm.Engine.Where("email = ? OR phone = ?", dto.Account, dto.Account).Get(&user)
	if err != nil {
		return nil, err
	}
	if has {
		return nil, resp.Error{Code: 100104, Msg: "user exist"}
	}

	success, err := u.redisService.ValidateCaptcha(ctx, dto.VerifyId, dto.Code)
	if err != nil {
		return nil, err
	}

	if success == false {
		return nil, resp.Error{Code: 110202, Msg: "CAPTCHA error"}
	}

	user = models.User{}

	if dto.Type == "email" {
		user.Email = dto.Account
	}
	if dto.Type == "phone" {
		user.Phone = dto.Account
	}

	_, err = orm.Engine.InsertOne(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil // 没有匹配的用户
}

func (u *AuthService) GetCaptcha(ctx context.Context, dto *dto.CaptchaGet) (string, string, error) {

	id := util.RandStringLetter(6)
	code := util.RandNumber(6)

	_, err := u.redisService.SetCaptcha(ctx, dto.Account, id, code)

	if err != nil {
		return "", "", err
	}

	return id, code, nil
}
