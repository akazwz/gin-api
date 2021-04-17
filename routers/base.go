package routers

import (
	v1 "github.com/akaedison/go-gin-demo/api/v1"
	"github.com/gin-gonic/gin"
)

func InitBaseRouter(Router *gin.RouterGroup) {
	Router.POST("/token", v1.CreateToken)
	Router.POST("/user", v1.CreateUser)
}
