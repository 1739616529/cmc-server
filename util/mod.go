package util

import (
	"fmt"
	"regexp"

	beego "github.com/beego/beego/v2/server/web"
)

func PrintApiPath() {
	tree := beego.PrintTree()
	methods := tree["Data"].(beego.M)
	for k, v := range methods {
		fmt.Printf("%s => %v\n", k, v)
	}
}

func VaildateEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

func VaildatePhone(phone string) bool {
	re := regexp.MustCompile(`^(\+?86)?-?1[3-9]\d{9}$`)
	return re.MatchString(phone)
}
