package routers

import (
	v1 "github.com/akazwz/go-gin-restful-api/api/v1"
	"github.com/gin-gonic/gin"
)

func InitBaseRouter(Router *gin.RouterGroup) {
	Router.POST("/token/username-pwd", v1.CreateTokenByUsernamePwd)
	Router.POST("/token/phone-pwd", v1.CreateTokenByPhonePwd)
	Router.POST("/users", v1.CreateUser)
	Router.POST("/avatar", v1.CreateFile)
}
