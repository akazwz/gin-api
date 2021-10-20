package routers

import (
	v1 "github.com/akazwz/go-gin-restful-api/api/v1"
	"github.com/gin-gonic/gin"
)

func InitVerificationCodeRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("verification")
	{
		UserRouter.GET("/sms", v1.GetVerificationCode)
		UserRouter.GET("/status", v1.GetVerificationStatus)
	}
}
