package initialize

import (
	"github.com/akaedison/go-gin-demo/routers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	var router = gin.Default()
	router.Use(cors.Default())
	group := router.Group("")
	{
		routers.InitUserRouter(group)
	}

	return router
}
