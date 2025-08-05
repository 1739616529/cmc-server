package rbac

import (
	"cmc-server/components/jwt"
	"cmc-server/components/logger"
	"cmc-server/components/orm"
	"cmc-server/models"
	"math/big"
	"strings"

	"github.com/beego/beego/v2/server/web/context"
)

var (
	Promisson map[string]string
)

func Init() {

	if err := InitRbacData(); err != nil {
		logger.Logger.Error("同步rbac失败: %v", err)
	}

	if err := CachePromission(); err != nil {
		logger.Logger.Error("缓存权限失败: %v", err)
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

func MatchPromission(promissionList string, promission string) bool {
	_promissionList := new(big.Int)
	_promission := new(big.Int)
	_result := new(big.Int)
	_promissionList.SetString(promissionList, 2)
	_promission.SetString(promission, 2)
	_result.And(_promissionList, _promission)
	return _result.Sign() != 0
}

func CachePromission() error {
	var userRole []models.UserRole

	if err := orm.Engine.Find(&userRole); err != nil {
		return err
	}

	for _, v := range userRole {
		var rolePromission models.RolePromission
		_, err := orm.Engine.Where("role_id = ?", v.RoleId).Get(&rolePromission)  

		if err != nil {
			return err
		}
		Promisson["role:"+v.UserId] = v.RoleId
		Promisson["promission"+v.RoleId] = rolePromission.RoleId
	}

	return nil
}
