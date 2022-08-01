package request

import (
	"github.com/golang-jwt/jwt/v4"
	uuid "github.com/satori/go.uuid"
)

type CustomClaims struct {
	BaseClaims
	BufferTime int64
	jwt.RegisteredClaims
}

type BaseClaims struct {
	UUID uuid.UUID
}
