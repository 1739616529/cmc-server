package common

import (
	"encoding/json"

	"github.com/beego/beego/v2/core/validation"
	"github.com/beego/beego/v2/server/web"
)

type BaseController struct {
	web.Controller
}

func (c *BaseController) ParseJson(dto any) error {
	return json.Unmarshal(c.Ctx.Input.RequestBody, dto)
}

func (c *BaseController) Send(res any) {

	c.Data["json"] = map[string]any{
		"code":    200,
		"message": "请求成功",
		"data":    res,
	}
	c.ServeJSON()
}

func (c *BaseController) VaildateError(msg string) {

	c.Ctx.Output.SetStatus(400)
	c.Data["json"] = map[string]any{
		"code":    400,
		"message": "error",
		"data":    msg,
	}
	c.ServeJSON()
}

func (c *BaseController) Vaildate(res any) bool {
	valid := validation.Validation{}
	passed, err := valid.Valid(res)
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Ctx.Output.Body([]byte("校验异常"))
		return false
	}
	if !passed {
		// 校验不通过，收集错误信息
		var errMsg string
		for _, err := range valid.Errors {
			errMsg += err.Key + "\n"
		}
		c.VaildateError(errMsg)
		return false
	}

	return true
}

func (c *BaseController) Error(err error) {

	c.Data["json"] = map[string]any{
		"code":    -1,
		"message": "error",
		"data":    err.Error(),
	}
	c.ServeJSON()
}
