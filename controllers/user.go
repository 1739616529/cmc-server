package controllers

import (
	"cmc-server/common"
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
