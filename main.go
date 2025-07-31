package main

import (
	"cmc-server/components/redis"
	_ "cmc-server/filter"
	"cmc-server/models"
	_ "cmc-server/routers"
	"cmc-server/util"

	beego "github.com/beego/beego/v2/server/web"
)

func main() {

	beego.BConfig.CopyRequestBody = true

	util.PrintApiPath()
	models.Init()
	redis.Init()

	beego.Run()
}
