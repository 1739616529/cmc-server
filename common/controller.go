package common

import (
	"cmc-server/components/logger"
	"cmc-server/resp"
	"encoding/json"
	"errors"

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
		"code":    0,
		"message": "success",
		"data":    res,
	}
	c.ServeJSON()
}

func (c *BaseController) VaildateError(msg string) {
	logger.Logger.Error("serverError: %#v", errors.New(msg))
	c.Error(400, errors.New(msg))
}

func (c *BaseController) Vaildate(res any) bool {
	valid := validation.Validation{}
	passed, err := valid.Valid(res)
	if err != nil {
		c.VaildateError(err.Error())
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

func (c *BaseController) Error(code int, err error) {

	if e, ok := err.(resp.Error); ok {
		c.Data["json"] = map[string]any{
			"code":    e.Code,
			"message": e.Msg,
			"data":    nil,
		}
	} else {
		c.Data["json"] = map[string]any{
			"code":    code,
			"message": err.Error(),
			"data":    nil,
		}
	}

	c.ServeJSON()
}

func (c *BaseController) ErrorMessage(code int, msg string) {

	c.Error(code, errors.New(msg))
}

func (c *BaseController) ServerError(err error) {
	logger.Logger.Error("serverError: %#v", err)
	c.Error(500, err)
}
