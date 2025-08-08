package controllers

import (
	"cmc-server/common"
	"cmc-server/dto"
	"cmc-server/service"
)

type PromissionController struct {
	common.BaseController
	promissionService service.PromissionService
}

func (c *PromissionController) ChangeRolePromission() {
	var req dto.RolePromissionChange

	c.ParseJson(&req)
	if ok := c.Vaildate(&req); !ok {
		return
	}

	c.promissionService.SetPromission(req)

}
