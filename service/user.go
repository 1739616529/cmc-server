package service

import (
	"cmc-server/components/orm"
	"cmc-server/models"
	"cmc-server/resp"
)

type UserService struct {
}

func (*UserService) FindUser(id string) (*models.User, error) {

	var user models.User

	has, err := orm.Engine.Where("id = ?", id).Get(&user)

	if err != nil {
		return nil, err
	}

	if !has {
		return nil, resp.NewError(resp.StatusUserNotFound)
	}

	return &user, nil

}
