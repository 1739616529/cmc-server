package routers

import (
	"cmc-server/components/jwt"
	"cmc-server/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func V1Router() *beego.Namespace {

	routes := []beego.LinkNamespace{}

	routes = append(routes, ControllerMethods("/user", &controllers.UserController{}, "/info:post:GetUser")...)

	routes = append(routes, ControllerMethods("/auth",
		&controllers.AuthController{},
		jwt.NoAuthPathPrefix+"/login:post:Login",
		jwt.NoAuthPathPrefix+"/register:post:Register",
		jwt.NoAuthPathPrefix+"/captcha:post:GetCaptcha",
	)...)

	return beego.NewNamespace("/api/v1",
		routes...,
	)
}
