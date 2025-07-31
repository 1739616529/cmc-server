package controllers

import (
	"cmc-server/common"
	"cmc-server/components/jwt"
	"cmc-server/dto"
	"cmc-server/models"
	"cmc-server/service"
)

type UserController struct {
	common.BaseController
	userService service.UserService
}

func (c *UserController) GetUser() {
	user := models.User{}
	user.Email = "astaxie@gmail.com"
	user.Name = "Job"
	user.Phone = "+01133336482889"
	user.Passwd = "aaaccc"
	c.Send(user.Output())
}

func (c *UserController) Login() {
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

	user, err := c.userService.UserLogin(&req)

	if err != nil {
		c.Error(err)
		return
	}

	token, err := jwt.Generate_token(user.Id)

	if err != nil {
		c.Error(err)
		return

	}

	c.Send(token)
}

func (c *UserController) Register() {
	var req dto.UserRegister
	c.ParseJson(&req)
	if ok := c.Vaildate(&req); !ok {
		return
	}
	// 账号或者邮箱为空
	if len(req.Email) == 0 && len(req.Phone) == 0 {
		c.VaildateError("Email or Phone Required.:")
		return
	}

	user, err := c.userService.UserRegister(&req)

	if err != nil {
		c.Error(err)
		return
	}

	token, err := jwt.Generate_token(user.Id)

	if err != nil {
		c.Error(err)
		return

	}

	c.Send(token)
}
