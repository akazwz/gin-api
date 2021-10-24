package routers

import (
	v1 "github.com/akazwz/go-gin-restful-api/api/v1"
	"github.com/gin-gonic/gin"
)

func InitSubRouter(Router *gin.RouterGroup) {
	subsRouter := Router.Group("subs")
	{
		subsRouter.GET("", v1.GetUserSub)
		subsRouter.POST("", v1.CreateSub)
		subsRouter.DELETE("", v1.DeleteSub)
	}
}
