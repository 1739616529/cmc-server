package controllers

import (
	"cmc-server/common"
	"cmc-server/components/jwt"
	"cmc-server/service"
)

type UserController struct {
	common.BaseController
	userService service.UserService
}

func (c *UserController) GetUser() {
	jwtPayload := c.Ctx.Input.GetData(jwt.JwtDataPayload).(*jwt.JwtPayload)

	user, err := c.userService.FindUser(jwtPayload.Id)

	if err != nil {
		c.ServerError(err)
		return
	}

	c.Send(user.Output())

}
