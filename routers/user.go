package routers

import (
	v1 "github.com/akazwz/go-gin-restful-api/api/v1"
	"github.com/akazwz/go-gin-restful-api/middleware"
	"github.com/gin-gonic/gin"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("users")
	{
		UserRouter.PATCH("/password", v1.ChangePassword)
		UserRouter.PATCH("/password/phone-code", v1.ChangePasswordByPhoneVerificationCode)
		UserRouter.PATCH("/profile", v1.UpdateUserProfile)
		UserRouter.Use(middleware.JWTAuthority777()).PATCH("/authority", v1.SetUserAuthority)
		UserRouter.Use(middleware.JWTAuthority777()).GET("", v1.GetUserList)
	}
}
