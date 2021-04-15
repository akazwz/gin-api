package initialize

import (
	"github.com/akaedison/go-gin-demo/middleware"
	"github.com/akaedison/go-gin-demo/routers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	var router = gin.Default()
	router.Use(cors.Default())

	publicRouter := router.Group("")
	{
		routers.InitBaseRouter(publicRouter)
	}

	privateGroup := router.Group("")
	privateGroup.Use(middleware.JWTAuth())
	{
		routers.InitUserRouter(privateGroup)
		routers.InitBookRouter(privateGroup)
	}

	return router
}
