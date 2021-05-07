package routers

import (
	v1 "github.com/akazwz/go-gin-restful-api/api/v1"
	"github.com/gin-gonic/gin"
)

func InitFileRouter(Router *gin.RouterGroup) {
	BookRouter := Router.Group("file")
	{
		BookRouter.POST("", v1.CreateFile)
		BookRouter.GET("status", v1.CreateFile)
	}
}
