package models

import (
	"cmc-server/common"
)

type RolePromission struct {
	common.BaseEntry `xorm:"extends"`
	UserId           string `xorm:"varchar(50) unique comment('用户id')"`
	RoleId           string `xorm:"varchar(20) unique comment('劫色id')"`
}
