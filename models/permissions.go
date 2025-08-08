package models

import (
	"cmc-server/common"
)

type Promission struct {
	common.BaseEntry `xorm:"extends"`
	Name             string `json:"name" xorm:"varchar(100) comment('名称')"`
	Code             string `json:"code" xorm:"varchar(100) unique comment('唯一代码')"`
	Path             string `json:"path" xorm:"varchar(100) unique comment('路由')"`
	Method           string `json:"method" xorm:"varchar(20) unique comment('请求方法')"`
	Description      string `json:"description" xorm:"varchar(20) default('') comment('描述') "`
	Rank             string `json:"rank" xorm:"varchar(1000) default('') comment('位阶') "`
}
