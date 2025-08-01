package logger

import (
	_ "cmc-server/env"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
)

var Logger *logs.BeeLogger

func init() {

	println("run mode: ", web.BConfig.RunMode)

	Logger = logs.NewLogger(1000)
	// 设置同时写到控制台和文件
	Logger.SetLogger(logs.AdapterFile, `{"filename":"logs/app.log","daily":true,"maxdays":7}`)

	// 设置日志级别（建议默认 info）
	Logger.SetLevel(logs.LevelInfo)

	if web.BConfig.RunMode == "dev" {
		Logger.SetLevel(logs.LevelDebug)
	} else {
		Logger.SetLevel(logs.LevelError)
	}

	// 开启异步（推荐）
	Logger.Async()

	// 输出调用的文件名和行号
	Logger.EnableFuncCallDepth(true)
}
