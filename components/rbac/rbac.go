package rbac

import (
	"cmc-server/components/orm"
	"cmc-server/models"
	"encoding/json"
	"errors"
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

	// 初始化 admin
	if _, err := InitAdminUser(); err != nil {
		return err
	}
	// 初始化 admin user role
	if err := InitUserRole(); err != nil {
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

		// 内置角色
		v.IsBuiltIn = true

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

func InitUserRole() error {

	// 1. 先看一下  admin user role 是否存在
	var userRole models.UserRole
	hasAdminUserRole, err := orm.Engine.Where("user_id = ?", "admin").Get(&userRole)
	if err != nil {
		return errors.New("get admin user role error:" + err.Error())
	}
	if hasAdminUserRole {
		return nil
	}

	// 2. 如果 admin user role 不存在，则创建 admin user role
	var adminRole models.Role
	hasAdminRole, err := orm.Engine.Where("code = ?", "ADMIN").Get(&adminRole)
	if err != nil {
		return errors.New("get admin role error:" + err.Error())
	}
	if !hasAdminRole {
		return errors.New("admin role not found")
	}
	userRole.UserId = "admin"
	userRole.RoleId = adminRole.Id
	_, err = orm.Engine.Insert(&userRole)

	if err != nil {
		return errors.New("inset admin user role error:" + err.Error())
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
