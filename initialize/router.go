package initialize

import (
	"github.com/akaedison/go-gin-demo/middleware"
	"github.com/akaedison/go-gin-demo/routers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Routers() *gin.Engine {
	var router = gin.Default()
	router.Use(cors.Default())

	router.Any("teapot", func(c *gin.Context) {
		c.JSON(http.StatusTeapot, gin.H{
			"message": "I'm a teapot",
			"story": "This code was defined in 1998 " +
				"as one of the traditional IETF April Fools' jokes," +
				" in RFC 2324, Hyper Text Coffee Pot Control Protocol," +
				" and is not expected to be implemented by actual HTTP servers." +
				" However, known implementations do exist.",
		})
	})

	publicRouterV1 := router.Group("v1")
	{
		routers.InitBaseRouter(publicRouterV1)
	}

	privateGroupV1 := router.Group("v1")
	privateGroupV1.Use(middleware.JWTAuth())
	{
		routers.InitUserRouter(privateGroupV1)
		routers.InitBookRouter(privateGroupV1)
	}

	return router
}
