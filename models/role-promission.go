package models

import (
	"cmc-server/common"
)

type RolePromission struct {
	common.BaseEntry `xorm:"extends"`
	RoleId           string `xorm:"varchar(20) comment('角色id')"`
	Promission       string `xorm:"varchar(20) defaut('0') comment('权限')"`
}
