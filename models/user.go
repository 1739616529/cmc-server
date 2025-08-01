package models

import (
	"cmc-server/common"
	"cmc-server/dto"
	"time"

	"github.com/google/uuid"
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
	Id        string    `xorm:"pk varchar(36)"`
	Name      string    `xorm:"varchar(50)"`
	Email     string    `xorm:"varchar(100) unique"`
	Phone     string    `xorm:"varchar(20) unique"`
	Passwd    string    `xorm:"varchar(255) notnull"`
	Status    int       `xorm:"notnull default(0)"`         // 状态
	Avatar    string    `xorm:"varchar(255)"`               // 头像地址
	Role      string    `xorm:"varchar(20) default 'user'"` // 角色字段
	IsDeleted bool      `xorm:"default false"`              // 软删除标记（可搭配逻辑删除实现）
	CreatedAt time.Time `xorm:"created"`                    // 创建时间（自动填充）
	UpdatedAt time.Time `xorm:"updated"`                    // 更新时间（自动更新时间）
	LastLogin time.Time `xorm:""`                           // 最近登录时间（可手动更新）
}

var _ common.Output = (*User)(nil)

// 在插入前设置 UUID
func (u *User) BeforeInsert() {
	println("BeforeInsert")
	if u.Id == "" {
		u.Id = uuid.New().String()
	}
}

func (u *User) Output() any {
	var userOV dto.UserOutput
	_ = copier.Copy(&userOV, u)
	return userOV
}
