package orm

import (
	"cmc-server/models"

	"github.com/beego/beego/v2/core/logs"
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

var Engine *xorm.Engine

func Init() {
	var err error
	dsn := "admin_dys:rLz48TThKQ8McCdk@tcp(mysql2.sqlpub.com:3307)/cursor_pool?charset=utf8"
	Engine, err = xorm.NewEngine("mysql", dsn)
	if err != nil {
		logs.Error("数据库连接失败: %v", err)
	}

	// 打印 SQL 语句（调试用）
	// Engine.ShowSQL(true)

	// 自动同步结构体到数据库（建表）
	err = Engine.Sync2(
		new(models.User),
		new(models.Promission),
		new(models.RolePromission),
		new(models.Role),
		new(models.UserRole),
	)

	if err != nil {
		logs.Error("同步数据库结构失败: %v", err)
	}

	err = InitRbacData()

	if err != nil {
		logs.Error("同步rbac 失败: %v", err)
	}
}
