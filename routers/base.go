package routers

import (
	v1 "github.com/akaedison/go-gin-demo/api/v1"
	"github.com/gin-gonic/gin"
)

func InitBaseRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("base")
	{
		UserRouter.POST("/login", v1.Login)
	}
}
