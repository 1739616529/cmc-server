package models

import (
	"cmc-server/common"
)

type Role struct {
	common.BaseEntry `xorm:"extends"`
	Name             string `json:"name" xorm:"varchar(100) comment('名称')"`
	Code             string `json:"code" xorm:"varchar(100) unique comment('唯一代码')"`
	Description      string `json:"description" xorm:"varchar(1000) comment('描述') default('')"`
	Promission       string `json:"promission" xorm:"varchar(1000) comment('权限') default('')"`
	UserId           string `json:"userId" xorm:"varchar(36) comment('用户ID')"`
}
