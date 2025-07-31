package models

import (
	"cmc-server/common"
	"cmc-server/dto"

	"github.com/google/uuid"
)

type User struct {
	Id     string `xorm:"pk varchar(36)"`
	Name   string `xorm:"varchar(100)"`
	Email  string `xorm:"varchar(100) unique"`
	Phone  string `xorm:"varchar(20) unique"`
	Passwd string `xorm:"varchar(100)"`
}

var _ common.Output = (*User)(nil)

// 在插入前设置 UUID
func (u *User) BeforeInsert() {
	if u.Id == "" {
		u.Id = uuid.New().String()
	}
}

func (u *User) Output() any {
	output := dto.UserOutput{}
	output.Email = u.Email
	output.Name = u.Name
	output.Phone = u.Phone
	return output
}
