package service

import (
	"cmc-server/components/redis"
	"cmc-server/dto"
	"cmc-server/models"
	"cmc-server/util"
	"context"
	"errors"
)

type AuthService struct {
	redisService redis.RedisService
}

func (u *AuthService) Login(dto *dto.UserLogin) (*models.User, error) {

	var user models.User

	has, err := models.Engine.Where("email = ? AND passwd = ?", dto.Email, dto.Passwd).Get(&user)
	if err != nil {
		return nil, err
	}
	if has {
		return &user, nil
	}

	// 邮箱没找到，再用手机号匹配
	has, err = models.Engine.Where("phone = ? AND passwd = ?", dto.Phone, dto.Passwd).Get(&user)
	if err != nil {
		return nil, err
	}
	if has {
		return &user, nil
	}

	return nil, nil // 没有匹配的用户
}

func (u *AuthService) Register(dto *dto.UserRegister) (*models.User, error) {

	var user models.User
	has, err := models.Engine.Where("email = ? OR phone = ?", dto.Email, dto.Phone).Get(&user)
	if err != nil {
		return nil, err
	}
	if has {
		return nil, errors.New("user is exist")
	}

	return &models.User{Id: "321"}, nil // 没有匹配的用户
}

func (u *AuthService) GetCaptcha(ctx context.Context, dto *dto.CaptchaGet) (string, error) {

	id := util.RandString(6)

	_, err := u.redisService.SetCaptcha(ctx, dto.Account, id, util.RandString(6))

	if err != nil {
		return "", err
	}

	return id, nil
}
