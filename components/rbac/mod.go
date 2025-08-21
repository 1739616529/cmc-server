package rbac

import (
	"cmc-server/components/jwt"
	"cmc-server/components/logger"
	"cmc-server/components/orm"
	"cmc-server/models"
	"math/big"
	"net/http"
	"strings"

	"github.com/beego/beego/v2/server/web/context"
)

var (
	Promisson map[string]string = make(map[string]string)
)

func Init() {

	if err := InitRbacData(); err != nil {
		logger.Logger.Error("同步rbac失败: %v", err)
	}

}

func RbacFilter(ctx *context.Context) {

	// 如果不走 jwt 就跳过权限鉴权
	path := ctx.Input.URL()
	if strings.Contains(path, jwt.NoAuthPathPrefix) {
		return
	}

	// 用户 id
	userId := ctx.Input.GetData(jwt.JwtDataPayload).(*jwt.JwtPayload).Id

	// 如果是admin 不鉴权
	if userId == "admin" {
		return
	}

	// 查找当前接口的权限
	routerPattern := ctx.Input.GetData("routerPattern").(string)
	var promission models.Promission
	hasPromission, err := orm.Engine.Where("? LIKE CONCAT('%', path)", routerPattern).Get(&promission)

	if err != nil {
		ctx.Output.SetStatus(http.StatusInternalServerError)
		ctx.Output.Body([]byte("rbac error. query promission error:" + err.Error()))
		return
	}

	// 如果没找到 说明不需要控权限
	if !hasPromission {
		return
	}

	isPromission, err := orm.Engine.SQL(`
        SELECT 1
        FROM user_role ur
        JOIN role_permission rp ON ur.role_id = rp.role_id
        JOIN permission p ON (rp.permission & p.bit) > 0
        WHERE ur.user_id = ?
            AND ? LIKE CONCAT('%', p.path)
        LIMIT 1;`, userId, routerPattern,
	).Count()

	if err != nil {
		ctx.Output.SetStatus(http.StatusInternalServerError)
		ctx.Output.Body([]byte("rbac error. query promission error:" + err.Error()))
		return
	}

	if isPromission == 0 {
		ctx.Output.SetStatus(http.StatusForbidden)
		ctx.Output.Body([]byte("Permission denied"))
		return
	}

}

func MatchPromission(promissionList string, promission string) bool {
	_promissionList := new(big.Int)
	_promission := new(big.Int)
	_result := new(big.Int)
	_promissionList.SetString(promissionList, 2)
	_promission.SetString(promission, 2)
	_result.And(_promissionList, _promission)
	return _result.Sign() != 0
}
