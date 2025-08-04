package models

import (
	"cmc-server/common"
)

type RolePromission struct {
	common.BaseEntry `xorm:"extends"`
	UserId           string `xorm:"varchar(50) comment('用户id')"`
	RoleId           string `xorm:"varchar(20) comment('角色id')"`
}
