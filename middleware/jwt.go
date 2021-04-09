package middleware

import (
	"errors"
	"github.com/akaedison/go-gin-demo/global"
	"github.com/akaedison/go-gin-demo/model/request"
	"github.com/akaedison/go-gin-demo/model/response"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		//jwt鉴权取头部信息 x-token,登录时返回token信息 前端进行存储
		token := c.Request.Header.Get("x-token")
		if token == "" {
			response.FailWithMessage("Not Login", c)
			c.Abort()
			return
		}

	}
}

type JWT struct {
	SigningKey []byte
}

var (
	TokenExpired = errors.New("token is expired")
)

func NewJWT() *JWT {
	return &JWT{
		SigningKey: []byte(global.CFG.JWT.SigningKey),
	}
}

func (j *JWT) ParseToken(tokenString string) (*request.CustomClaims, error) {
	_, err := jwt.ParseWithClaims(tokenString, &request.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})

	if err != nil {

	}

	return nil, err
}
