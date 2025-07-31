package routers

import (
	"cmc-server/components/jwt"
	"cmc-server/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func V1Router() *beego.Namespace {
	return beego.NewNamespace("/api/v1",
		ControllerMethods("/user",
			&controllers.UserController{},
			"/get:post:GetUser",
			jwt.NoAuthPathPrefix+"/login:post:Login",
			jwt.NoAuthPathPrefix+"/register:post:Register",
		)...,
	)
}
