package models

import (
	"cmc-server/common"
	"cmc-server/dto"
	"time"

	"github.com/jinzhu/copier"
)

/*
*
| 值 | 状态名     | 含义说明                  |
| - | ---------  | ---------------------     |
| 0 | 正常       | 默认状态，用户处于正常可登录/使用状态   |
| 1 | 禁用       | 用户被管理员封禁，无法登录或使用服务    |
| 2 | 审核中     | 用户注册后进入审核流程，尚未开放权限    |
| 3 | 审核未通过  | 用户审核失败，可能需修改资料后重审     |
| 4 | 注销       | 用户主动注销或被管理员注销（逻辑关闭账户） |
| 5 | 冻结       | 因异常行为或风控限制，暂时无法使用账号   |
*/
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
