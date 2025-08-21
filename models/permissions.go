package models

import (
	"cmc-server/common"
	"cmc-server/dto"

	"github.com/jinzhu/copier"
)

type Promission struct {
	common.BaseEntry `xorm:"extends"`
	Name             string `json:"name" xorm:"varchar(100) notnull comment('权限名称')"`
	Code             string `json:"code" xorm:"varchar(100) unique notnull comment('权限唯一代码')"`
	Path             string `json:"path" xorm:"varchar(200) comment('路由路径')"`
	Method           string `json:"method" xorm:"varchar(20) comment('请求方法')"`
	Description      string `json:"description" xorm:"varchar(500) default('') comment('权限描述')"`
	Bit              string `json:"bit" xorm:"varchar(100) default('') comment('权限级别')"`
}

var _ common.Output = (*Promission)(nil)

func (p *Promission) Output() any {
	var permissionOutput dto.PermissionOutput
	_ = copier.Copy(&permissionOutput, p)
	return permissionOutput
}
