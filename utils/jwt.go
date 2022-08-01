package utils

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/akazwz/gin-api/model/request"
	"github.com/golang-jwt/jwt/v4"
)

type JWT struct {
	SigningKey []byte
}

var (
	TokenExpired     = errors.New("token is expired")
	TokenNotValidYet = errors.New("token not active yet")
	TokenMalformed   = errors.New("that's not even a token")
	TokenInvalid     = errors.New("couldn't handle this token")
)

func NewJWT() *JWT {
	return &JWT{SigningKey: []byte(os.Getenv("JWT_SIGNING_KEY"))}
}

func (j *JWT) CreateClaims(baseClaims request.BaseClaims) request.CustomClaims {
	bufferTime, _ := strconv.ParseInt(os.Getenv("JWT_BUFFER_TIME"), 10, 64)
	hours, _ := strconv.Atoi(os.Getenv("JWT_EXPIRES_AT"))
	claims := request.CustomClaims{
		BaseClaims: baseClaims,
		BufferTime: bufferTime,
		/* StandardClaims 已废弃,改用 RegisterClaims */
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: os.Getenv("JWT_ISSUER"),
			/* int 需要转换为 time.Duration  */
			ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(time.Duration(hours) * time.Hour)},
			NotBefore: &jwt.NumericDate{Time: time.Now()},
		},
	}
	return claims
}

// CreateToken 创建token
func (j *JWT) CreateToken(claims request.CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// ParseToken 解析 token
func (j *JWT) ParseToken(tokenStr string) (*request.CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &request.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		ve, ok := err.(*jwt.ValidationError)
		if ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
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
