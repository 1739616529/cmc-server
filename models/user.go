package models

import (
	"cmc-server/common"
	"cmc-server/dto"
	"time"

	"github.com/jinzhu/copier"
)

type User struct {
	common.BaseEntry `xorm:"extends"`
	Name             string    `xorm:"varchar(50) comment('名称')"`
	Email            string    `xorm:"varchar(100) unique comment('邮箱')"`
	Phone            string    `xorm:"varchar(20) unique comment('手机号')"`
	Passwd           string    `xorm:"varchar(255) notnull comment('密码')"`
	Avatar           string    `xorm:"varchar(255) comment('头像地址')"`
	LastLogin        time.Time `xorm:"comment('最近登录时间')"`
}

var _ common.Output = (*User)(nil)

func (u *User) Output() any {
	var userOV dto.UserOutput
	_ = copier.Copy(&userOV, u)
	return userOV
}
