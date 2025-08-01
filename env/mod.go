package env

import (
	"os"

	"github.com/beego/beego/v2/server/web"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load(".env")
	web.BConfig.RunMode = os.Getenv("APP_ENV")
}
