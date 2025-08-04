package models

import (
	"cmc-server/common"
)

type Role struct {
	common.BaseEntry `xorm:"extends"`
	Name             string `xorm:"varchar(50) comment('名称')"`
	Code             string `xorm:"varchar(20) unique comment('唯一代码')"`
	Description      string `xorm:"varchar(20) comment('描述') default('')"`
}
