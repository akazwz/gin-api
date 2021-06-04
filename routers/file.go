package routers

import (
	v1 "github.com/akazwz/go-gin-restful-api/api/v1"
	"github.com/gin-gonic/gin"
)

func InitFileRouter(Router *gin.RouterGroup) {
	FileRoute := Router.Group("file")
	{
		FileRoute.POST("", v1.CreateFile)
	}
}
