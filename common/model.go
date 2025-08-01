package common

import (
	"cmc-server/util"
	"time"
)

type Output interface {
	Output() any
}

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
		u.Id = util.Guid.String()
	}
}
