package rbac

import (
	"cmc-server/components/orm"
	"cmc-server/models"
	"encoding/json"
	"os"
)

var (
	ADMIN_CODE = "ADMIN"
)

func InitRbacData() error {

	// 初始化权限
	if err := InitPromission(); err != nil {
		return err
	}

	// 初始化 角色
	if err := InitRole(); err != nil {
		return err
	}

	return nil
}

func InitPromission() error {

	_, err := orm.Engine.Exec("TRUNCATE TABLE promission")

	if err != nil {
		return err
	}

	data, err := os.ReadFile("static/promission.json")
	if err != nil {
		return err
	}

	var promissionList []models.Promission

	err = json.Unmarshal(data, &promissionList)

	if err != nil {
		return err
	}

	for _, v := range promissionList {
		_, err = orm.Engine.Insert(&v)
		if err != nil {
			return err
		}
	}

	return nil
}

func InitRole() error {

	data, err := os.ReadFile("static/role.json")
	if err != nil {
		return err
	}

	var roleList []models.Role

	err = json.Unmarshal(data, &roleList)

	if err != nil {
		return err
	}

	for _, v := range roleList {
		var existRole models.Role
		exists, err := orm.Engine.Where("code = ?", v.Code).Get(&existRole)
		if err != nil {
			return err
		}

		if exists {
			orm.Engine.ID(existRole.Id).Update(&v)
		} else {
			orm.Engine.Insert(&v)
		}
	}

	return nil
}

func InitAdminUser() (*models.User, error) {

	var user models.User
	adminUserHas, err := orm.Engine.Where("id = ?", "admin").Get(&user)
	if err != nil {
		return nil, err
	}

	if !adminUserHas {
		user = models.User{}
		user.Id = "admin"
		user.Email = "admin@qq.com"
		user.Passwd = "1.123."

		orm.Engine.InsertOne(&user)
	}

	return &user, nil

}
