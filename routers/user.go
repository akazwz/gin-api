package routers

import (
	v1 "github.com/akazwz/go-gin-demo/api/v1"
	"github.com/akazwz/go-gin-demo/middleware"
	"github.com/gin-gonic/gin"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("users")
	{
		UserRouter.PATCH("/password", v1.ChangePassword)
		UserRouter.Use(middleware.JWTAuthority777()).PATCH("/authority", v1.SetUserAuthority)
		UserRouter.Use(middleware.JWTAuthority777()).GET("", v1.GetUserList)
	}
}
