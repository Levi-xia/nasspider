package middler

import (
	"charging/config"
	"charging/internal/common"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// check登陆
func NeedLoginHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		tk_str, _ := c.Cookie("AT")
		if tk_str == "" {
			c.JSON(http.StatusOK, &common.Result{
				Code: 201,
				Msg:  "请先登录",
				Data: nil,
			})
			c.Abort()
			return
		}
		if token, err := tokenParse(tk_str); err != nil || !token.Valid {
			c.JSON(http.StatusOK, &common.Result{
				Code: 201,
				Msg:  "请先登录",
				Data: nil,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}

func tokenParse(tk_str string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(tk_str, &common.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return jwt.ErrInvalidKey, nil
		}
		return []byte(config.Conf.Jwt.Secret), nil
	})
	return token, err
}
