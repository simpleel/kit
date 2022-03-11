package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JwtAuth struct {
	key, secret   string
	expire, uniId int64
}

func NewAuth(key, secret string, expire, uniId int64) Auth {
	return &JwtAuth{
		key:    key,
		secret: secret,
		expire: expire,
		uniId:  uniId,
	}
}

// GetToken 获取或刷新 JWT Token
func (a *JwtAuth) GetToken() (now int64, jwtToken string, err error) {
	now = time.Now().Unix()

	claims := make(jwt.MapClaims)
	claims["exp"] = now + a.expire
	claims["iat"] = now
	claims[a.key] = fmt.Sprint(a.uniId)

	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	jwtToken, err = token.SignedString([]byte(a.secret))

	return
}

func GetSecretValue(ctx context.Context, key string) interface{} {
	return ctx.Value(key)
}
