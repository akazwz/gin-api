package routers

import (
	v1 "github.com/akazwz/go-gin-restful-api/api/v1"
	"github.com/gin-gonic/gin"
)

func InitBaseRouter(Router *gin.RouterGroup) {
	Router.POST("/token", v1.CreateToken)
	Router.POST("/users", v1.CreateUser)
	Router.POST("/avatar", v1.CreateFile)
}
