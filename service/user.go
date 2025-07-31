package service

import (
	"cmc-server/dto"
	"cmc-server/models"
	"errors"
)

type UserService struct{}

func (u *UserService) UserLogin(dto *dto.UserLogin) (*models.User, error) {

	var user models.User

	// 创建 xorm 引擎，假设已经初始化为 engine
	// engine, err := xorm.NewEngine("mysql", "user:pass@/dbname")
	// if err != nil {
	//     return nil, err
	// }

	// 先用邮箱匹配
	has, err := models.Engine.Where("email = ? AND password = ?", dto.Email, dto.Passwd).Get(&user)
	if err != nil {
		return nil, err
	}
	if has {
		return &user, nil
	}

	// 邮箱没找到，再用手机号匹配
	has, err = models.Engine.Where("phone = ? AND password = ?", dto.Phone, dto.Passwd).Get(&user)
	if err != nil {
		return nil, err
	}
	if has {
		return &user, nil
	}

	return nil, nil // 没有匹配的用户
}

func (u *UserService) UserRegister(dto *dto.UserRegister) (*models.User, error) {

	var user models.User

	// 创建 xorm 引擎，假设已经初始化为 engine
	// engine, err := xorm.NewEngine("mysql", "user:pass@/dbname")
	// if err != nil {
	//     return nil, err
	// }

	// 先用邮箱匹配
	has, err := models.Engine.Where("email = ? AND phone = ?", dto.Email, dto.Phone).Get(&user)
	if err != nil {
		return nil, err
	}
	if has {
		return nil, errors.New("user is exist")
	}

	return nil, nil // 没有匹配的用户
}
