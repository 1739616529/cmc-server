package jwt

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/beego/beego/v2/server/web/context"
	gm "github.com/tjfoc/gmsm/sm4"
)

var (
	Key              = []byte("1234567890abcdef")
	vi               = make([]byte, 8)
	JwtDataUserId    = "jwt.filter.userid"
	JwtDataPayload   = "jwt.filter.JwtPayload"
	NoAuthPathPrefix = "/noAuth"
)

func init() {
}

type JwtPayload struct {
	Id         string `json:"id"`
	Expiration int64  `json:"expiration"`
}

func Generate_token(id string) (string, error) {
	now := time.Now()
	future := now.AddDate(0, 0, 30)
	timestamp := future.UnixNano()
	payload := JwtPayload{Id: id, Expiration: timestamp}

	byt, err := JwtEncrypt(payload)
	if err != nil {
		return "", err
	}

	return string(*byt), nil
}

func JwtFilter(ctx *context.Context) {

	path := ctx.Input.URL()
	// 检查是否是排除路径
	if strings.Contains(path, NoAuthPathPrefix) {
		return // 放行，不执行拦截逻辑
	}
	authHeader := ctx.Input.Header("Authorization")
	if authHeader == "" {
		ctx.Output.SetStatus(http.StatusUnauthorized)
		ctx.Output.Body([]byte("Missing Authorization header"))
		return
	}

	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		ctx.Output.SetStatus(http.StatusUnauthorized)
		ctx.Output.Body([]byte("Invalid Authorization header format"))
		return
	}

	tokenString := parts[1]

	jwtPayload, err := JwtDecrypt([]byte(tokenString))

	if err != nil {
		ctx.Output.SetStatus(400)
		ctx.Output.Body([]byte("parse token error: " + err.Error()))
		return
	}

	timestramp := time.Now().UnixNano()

	if jwtPayload.Expiration < timestramp {
		ctx.Output.SetStatus(http.StatusUnauthorized)
		ctx.Output.Body([]byte("Please log in again after the token expires"))
		return
	}

	ctx.Input.SetData(JwtDataPayload, jwtPayload)
	ctx.Input.SetData(JwtDataUserId, jwtPayload.Id)

}

func JwtEncrypt(payload JwtPayload) (*[]byte, error) {
	var (
		err error
	)
	str, err := json.Marshal(payload)

	str = append(str, []byte("123")...)

	fmt.Println("aaaaaaa", len(str))
	if err != nil {
		return nil, err
	}

	byt, _, err := gm.Sm4GCM(Key, vi, str, []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef, 0xfe, 0xdc, 0xba, 0x98, 0x76, 0x54, 0x32, 0x10}, true)

	if err != nil {
		return nil, err
	}

	return &byt, nil
}

func JwtDecrypt(byt []byte) (*JwtPayload, error) {
	var (
		err error
	)
	json_str, _, err := gm.Sm4GCM(Key, vi, byt, []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef, 0xfe, 0xdc, 0xba, 0x98, 0x76, 0x54, 0x32, 0x10}, false)
	if err != nil {
		return nil, err
	}

	var payload JwtPayload

	if err := json.Unmarshal([]byte(json_str), &payload); err != nil {
		return nil, err
	}

	return &payload, nil

}

func init() {
	fmt.Println("aaa")
	aaa, _ := Generate_token("321")

	fmt.Println("data", aaa)

	json_str, _ := JwtDecrypt([]byte(aaa))
	fmt.Println("data2", json_str)
}
