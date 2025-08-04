package common

import (
	"time"

	"github.com/rs/xid"
)

type Output interface {
	Output() any
}

/*
*
| 值 | 状态名
| - |
| 0 | 正常
| 1 | 禁用
| 2 | 审核中
| 3 | 审核未通过
| 4 | 注销
| 5 | 冻结
*/

type BaseEntry struct {
	Id        string    `xorm:"pk varchar(36)"`
	IsDeleted bool      `xorm:"default false comment('软删除')"`
	CreatedAt time.Time `xorm:"created comment('创建时间')"`
	UpdatedAt time.Time `xorm:"updated comment('更新时间')"`
	Status    int       `xorm:"notnull default(0) comment('状态')"`
}

// 在插入前设置 UUID
func (u *BaseEntry) BeforeInsert() {

	if u.Id == "" {
		u.Id = xid.New().String()
	}
}

var _ Output = (*BaseEntry)(nil)

func (u *BaseEntry) Output() any {
	return *u
}
