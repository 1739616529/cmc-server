package controllers

import (
	"cmc-server/common"
	"cmc-server/components/jwt"
	"cmc-server/dto"
	"cmc-server/service"
	"errors"
)

type AuthController struct {
	common.BaseController
	authService service.AuthService
}

func (c *AuthController) Login() {
	var req dto.UserLogin
	c.ParseJson(&req)
	if ok := c.Vaildate(&req); !ok {
		return
	}
	// 账号或者邮箱为空
	if len(req.Email) == 0 && len(req.Phone) == 0 {
		c.VaildateError("Email or Phone Required.:")
		return
	}

	user, err := c.authService.Login(&req)

	if err != nil {
		c.ServerError(err)
		return
	}

	// 用户不存在
	if user == nil {
		c.Error(100104, errors.New("user not fount"))
		return
	}

	token, err := jwt.JwtEncrypt(user.Id)

	if err != nil {
		c.ServerError(err)
		return

	}

	c.Send(token)
}

func (c *AuthController) GetCaptcha() {
	var payload dto.CaptchaGet
	c.ParseJson(&payload)
	if ok := c.Vaildate(&payload); !ok {
		return
	}

	ctx := c.Ctx.Request.Context()

	id, code, err := c.authService.GetCaptcha(ctx, &payload)

	if err != nil {
		c.ServerError(err)
		return
	}

	c.Send(map[string]string{
		"id":   id,
		"code": code,
	})
}

func (c *AuthController) Register() {
	var req dto.UserRegister
	c.ParseJson(&req)

	if ok := c.Vaildate(&req); !ok {
		return
	}

	ctx := c.Ctx.Request.Context()

	user, err := c.authService.Register(ctx, &req)

	if err != nil {
		c.ServerError(err)
		return
	}

	token, err := jwt.JwtEncrypt(user.Id)

	if err != nil {
		c.ServerError(err)
		return
	}

	c.Send(token)
}
