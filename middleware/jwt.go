package middleware

import (
	"errors"
	"github.com/akazwz/go-gin-demo/global"
	"github.com/akazwz/go-gin-demo/model/request"
	"github.com/akazwz/go-gin-demo/model/response"
	"github.com/akazwz/go-gin-demo/service"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

const (
	CodeTokenNUll        = 4010
	CodeTokenExpired     = 4011
	CodeTokenNotValidYet = 4012
	CodeTokenMalformed   = 4013
	CodeTokenInvalid     = 4014
	CodeNoSuchUser       = 4015
	CodePermissionDenied = 4016
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			response.Unauthorized(CodeTokenNUll, "no token", c)
			c.Abort()
			return
		}

		j := NewJWT()
		claims, err := j.ParseToken(token)
		if err != nil {
			switch err {
			case TokenExpired:
				response.Unauthorized(CodeTokenExpired, "Token Expired", c)
				c.Abort()
				return
			case TokenNotValidYet:
				response.Unauthorized(CodeTokenNotValidYet, "Token Not Valid Yet", c)
				c.Abort()
				return
			case TokenMalformed:
				response.Unauthorized(CodeTokenMalformed, "Token Malformed", c)
				c.Abort()
				return
			case TokenInvalid:
				response.Unauthorized(CodeTokenInvalid, "Token Invalid", c)
				c.Abort()
				return
			default:
				c.Abort()
				return
			}
		}

		if err, _ = service.FindUserByUUID(claims.UUID.String()); err != nil {
			response.Unauthorized(CodeNoSuchUser, "No Such User", c)
			c.Abort()
			return
		}

		if claims.ExpiresAt-time.Now().Unix() < claims.BufferTime {
			claims.ExpiresAt = time.Now().Unix() + global.CFG.JWT.ExpiresTime
			newToken, _ := j.CreateToken(*claims)
			newClaims, _ := j.ParseToken(newToken)
			c.Header("'new-token", newToken)
			c.Header("'new-expires-at", strconv.FormatInt(newClaims.ExpiresAt, 10))
		}
		c.Set("claims", claims)
		c.Set("uuid", claims.UUID)
		c.Next()
	}
}

func JWTAuthority777() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		j := NewJWT()
		claims, _ := j.ParseToken(token)
		uuid := claims.UUID.String()
		_, user := service.FindUserByUUID(uuid)
		if user.AuthorityId != "777" {
			response.PermissionDenied(CodePermissionDenied, "Permission Denied", c)
			c.Abort()
			return
		}
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
