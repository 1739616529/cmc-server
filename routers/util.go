package routers

import (
	"strings"

	beego "github.com/beego/beego/v2/server/web"
	lo "github.com/samber/lo"
)

func ControllerMethods(prefixPath string, c beego.ControllerInterface, mappingMethods ...string) []beego.LinkNamespace {
	return lo.Map(mappingMethods, func(mappingMethod string, index int) beego.LinkNamespace {
		parts := strings.SplitN(mappingMethod, ":", 2)
		return beego.NSRouter(prefixPath+parts[0], c, parts[1])
	})
}
