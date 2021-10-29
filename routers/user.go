package routers

import (
	v1 "github.com/akazwz/go-gin-restful-api/api/v1"
	"github.com/akazwz/go-gin-restful-api/middleware"
	"github.com/gin-gonic/gin"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("users")
	{
		UserRouter.PUT("/password", v1.ChangePassword)
		UserRouter.PUT("/password/phone-code", v1.ChangePasswordByPhoneVerificationCode)
		UserRouter.PUT("/profile", v1.UpdateUserProfile)
		UserRouter.POST("/notify", v1.SetNotify)
		UserRouter.Use(middleware.JWTAuthority777()).PUT("/authority", v1.SetUserAuthority)
		UserRouter.Use(middleware.JWTAuthority777()).GET("", v1.GetUserList)
	}
}
