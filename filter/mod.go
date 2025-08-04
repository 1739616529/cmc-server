package filter

import (
	"cmc-server/components/jwt"
	"cmc-server/components/rbac"

	"github.com/beego/beego/v2/server/web"
)

func init() {
	web.InsertFilter("/*", web.BeforeRouter, jwt.JwtFilter)
	web.InsertFilter("/*", web.BeforeRouter, rbac.RbacFilter)
}
