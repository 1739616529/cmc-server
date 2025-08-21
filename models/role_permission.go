package models

import (
	"cmc-server/common"
)

type RolePermission struct {
	common.BaseEntry `xorm:"extends"`
	RoleId           string `json:"roleId" xorm:"varchar(36) notnull comment('角色ID') index(idx_role_permission_unique)"`
	Permission       string `json:"permission" xorm:"varchar(1000) notnull comment('权限') index(idx_role_permission_unique)"`
}

// 设置表名
func (RolePermission) TableName() string {
	return "role_permission"
}
