package models

import (
	"cmc-server/common"
	"cmc-server/dto"

	"github.com/jinzhu/copier"
)

type Role struct {
	common.BaseEntry `xorm:"extends"`
	Name             string `json:"name" xorm:"varchar(100) notnull comment('角色名称')"`
	Code             string `json:"code" xorm:"varchar(100) unique notnull comment('角色唯一代码')"`
	Description      string `json:"description" xorm:"varchar(500) comment('角色描述') default('')"`
	IsBuiltIn        bool   `json:"isBuiltIn" xorm:"default false comment('是否内置角色')"`
}

var _ common.Output = (*Role)(nil)

func (r *Role) Output() any {
	var roleOutput dto.RoleOutput
	_ = copier.Copy(&roleOutput, r)
	return roleOutput
}
