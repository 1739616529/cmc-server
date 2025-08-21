package service

import (
	"cmc-server/components/orm"
	"cmc-server/dto"
	"cmc-server/models"
)

type PromissionService struct {
}

func (*PromissionService) SetPromission(dto dto.RolePromissionChange) error {

	// var role models.Role

	// hasRolePromission, err := orm.Engine.Where("id = ?", dto.RoleID).Get(&rolePromission)

	// if err != nil {
	// 	return err
	// }

	// if !hasRolePromission {
	// 	return resp.NewError(resp.StatusRolePrimissionNotFound)
	// }

	// rolePromission.Promission = dto.Promission
	// rolePromission.Status = dto.Status

	// _, err = orm.Engine.ID(rolePromission.Id).Update(&rolePromission)
	// if err != nil {
	// 	return err
	// }

	return nil
}

func (*PromissionService) RolAdd(dto dto.RoleAdd) error {

	role := models.Role{
		Name:        dto.Name,
		Code:        dto.Code,
		Description: dto.Description,
	}

	_, err := orm.Engine.Insert(&role)

	if err != nil {
		return err
	}

	return nil
}
