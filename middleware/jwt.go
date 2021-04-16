package middleware

import (
	"errors"
	"github.com/akaedison/go-gin-demo/global"
	"github.com/akaedison/go-gin-demo/model/request"
	"github.com/akaedison/go-gin-demo/model/response"
	"github.com/akaedison/go-gin-demo/service"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		//jwt鉴权取头部信息 token,登录时返回token信息 前端进行存储
		token := c.Request.Header.Get("token")
		if token == "" {
			response.FailWithMessage("Not Login", c)
			c.Abort()
			return
		}
		j := NewJWT()

		claims, err := j.ParseToken(token)
		if err != nil {
			if err == TokenExpired {
				response.FailWithMessage("token expired", c)
				c.Abort()
				return
			}
			response.FailWithMessage(err.Error(), c)
			c.Abort()
			return
		}
		if err, _ = service.FindUserByUUID(claims.UUID.String()); err != nil {
			response.FailWithMessage(err.Error(), c)
			c.Abort()
		}
		if claims.ExpiresAt-time.Now().Unix() < claims.BufferTime {
			claims.ExpiresAt = time.Now().Unix() + global.CFG.JWT.ExpiresTime
			newToken, _ := j.CreateToken(*claims)
			newClaims, _ := j.ParseToken(newToken)
			c.Header("'new-token", newToken)
			c.Header("'new-expires-at", strconv.FormatInt(newClaims.ExpiresAt, 10))
		}
		c.Set("claims", claims)
		c.Next()
	}
}

type JWT struct {
	SigningKey []byte
}

var (
	TokenExpired     = errors.New("token is expired")
	TokenNotValidYet = errors.New("token not active yet")
	TokenMalformed   = errors.New("that`s not even a token")
	TokenInvalid     = errors.New("could`t handle this token")
)

func NewJWT() *JWT {
	return &JWT{
		SigningKey: []byte(global.CFG.JWT.SigningKey),
	}
}

func (j *JWT) CreateToken(claims request.CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

func (j *JWT) ParseToken(tokenString string) (*request.CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &request.CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}

	if token != nil {
		if claims, ok := token.Claims.(*request.CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid
	} else {
		return nil, TokenInvalid
	}

}
