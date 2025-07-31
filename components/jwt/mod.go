package jwt

import (
	"cmc-server/util"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/beego/beego/v2/server/web/context"
)

var (
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

	jwtPayload, err := JwtDecrypt(tokenString)

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

func JwtEncrypt(id string) (string, error) {
	now := time.Now()
	future := now.AddDate(0, 0, 30)
	timestamp := future.UnixNano()
	payload := JwtPayload{Id: id, Expiration: timestamp}

	jwt_json_str, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}
	byt, err := util.Encrypt(jwt_json_str)
	if err != nil {
		return "", err
	}

	return string(*byt), nil
}

func JwtDecrypt(token string) (*JwtPayload, error) {
	jwt_payload_str, err := util.Decrypt([]byte(token))
	if err != nil {
		return nil, err
	}

	var payload JwtPayload

	// 解析 JSON
	err = json.Unmarshal(*jwt_payload_str, &payload)
	if err != nil {
		return nil, err
	}

	return &payload, nil
}

func init() {
}
