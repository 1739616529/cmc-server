package models

import (
	"cmc-server/common"
)

type UserRole struct {
	common.BaseEntry `xorm:"extends"`
	UserId           string `json:"userId" xorm:"varchar(36) notnull comment('用户ID') index(idx_user_role_unique)"`
	RoleId           string `json:"roleId" xorm:"varchar(36) notnull comment('角色ID') index(idx_user_role_unique)"`
}

// 设置表名
func (UserRole) TableName() string {
	return "user_role"
}