package initialize

import (
	"github.com/akaedison/go-gin-demo/routers"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	var Router = gin.Default()
	group := Router.Group("")
	{
		routers.InitUserRouter(group)
	}

	return Router
}
