package service

import (
	"cmc-server/dto"
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
