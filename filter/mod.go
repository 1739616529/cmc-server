package filter

import (
	"cmc-server/components/jwt"

	"github.com/beego/beego/v2/server/web"
)

func init() {
	web.InsertFilter("/*", web.BeforeRouter, jwt.JwtFilter)
}
