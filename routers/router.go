package routers

import (
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.AddNamespace(
		V1Router(),
	)
}
