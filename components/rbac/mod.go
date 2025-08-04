package rbac

import (
	"cmc-server/components/jwt"
	"cmc-server/components/logger"
	"strings"

	"github.com/beego/beego/v2/server/web/context"
)

func Init() {
	err := InitRbacData()

	if err != nil {
		logger.Logger.Error("同步rbac 失败: %v", err)
	}
}

func RbacFilter(ctx *context.Context) {

	// 如果不走 jwt 就跳过权限鉴权
	path := ctx.Input.URL()
	if strings.Contains(path, jwt.NoAuthPathPrefix) {
		return
	}

	userId := ctx.Input.GetData(jwt.JwtDataPayload).(*jwt.JwtPayload).Id

	println("userId: ", userId)

}
