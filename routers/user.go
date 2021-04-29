package routers

import (
	v1 "github.com/akazwz/go-gin-demo/api/v1"
	"github.com/gin-gonic/gin"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("users")
	{
		UserRouter.PUT("/password", v1.ChangePassword)
		UserRouter.GET("", v1.GetUserList)
	}
}
