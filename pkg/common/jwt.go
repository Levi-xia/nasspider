package common

import (
	"nasspider/config"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JwtService struct {
}

type CustomClaims struct {
	jwt.StandardClaims
}

const (
	TokenType      = "bearer"
	AppGuardName   = "app"   // 端上
	AdminGuardName = "admin" // 后台
)

type TokenOutPut struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

func (jws *JwtService) CreateToken(GuardName string, userId string) (tokenData TokenOutPut, err error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		CustomClaims{
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Unix() + config.Conf.Jwt.JwtTtl,
				Id:        userId,
				Issuer:    GuardName, // 用于在中间件中区分不同客户端颁发的 token，避免 token 跨端使用
				NotBefore: time.Now().Unix() - 1000,
			},
		},
	)
	tokenStr, err := token.SignedString([]byte( config.Conf.Jwt.Secret))
	tokenData = TokenOutPut{
		tokenStr,
		int(config.Conf.Jwt.JwtTtl),
		TokenType,
	}
	return
}
