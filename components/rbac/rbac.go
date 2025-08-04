package rbac

import (
	"cmc-server/components/orm"
	"cmc-server/models"
	"encoding/json"
	"os"
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

	// 初始化 用户(admin) 角色
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
	_, err := orm.Engine.Exec("TRUNCATE TABLE role")

	if err != nil {
		return err
	}

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
		_, err = orm.Engine.Insert(&v)
		if err != nil {
			return err
		}
	}

	return nil
}

func InitUserRole() error {

	// 删除用户角色
	_, err := orm.Engine.Exec("TRUNCATE TABLE user_role")
	if err != nil {
		return err
	}

	// 角色权限
	_, err = orm.Engine.Exec("TRUNCATE TABLE role_promission")
	if err != nil {
		return err
	}

	// 获取 admin 用户
	adminUser, err := InitAdminUser()
	if err != nil {
		return err
	}

	// 获取 admin用户角色
	var userRole models.UserRole
	_, err = orm.Engine.Where("user_id = ?", adminUser.Id).Get(&userRole)

	if err != nil {
		return err
	}

	// 获取 admin 角色
	var adminRole models.Role
	_, err = orm.Engine.Where("code = ?", "ADMIN").Get(&adminRole)
	if err != nil {
		return err
	}

	// 创建 admin 角色权限
	userRole = models.UserRole{
		UserId: adminUser.Id,
		RoleId: adminRole.Id,
	}
	orm.Engine.Insert(&userRole)

	return nil
}

func InitAdminUser() (*models.User, error) {

	var user models.User
	adminUserHas, err := orm.Engine.Where("id = ?", "admin").Get(&user)
	if err != nil {
		return nil, err
	}

	if adminUserHas == false {
		user = models.User{}
		user.Id = "admin"
		user.Email = "admin@qq.com"
		user.Passwd = "1.123."

		orm.Engine.InsertOne(&user)
	}

	return &user, nil

}
