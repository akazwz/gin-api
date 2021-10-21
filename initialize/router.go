package initialize

import (
	"net/http"
	"time"

	"github.com/akazwz/go-gin-restful-api/api"
	"github.com/akazwz/go-gin-restful-api/middleware"
	"github.com/akazwz/go-gin-restful-api/model/response"
	"github.com/akazwz/go-gin-restful-api/routers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Routers() *gin.Engine {
	var router = gin.Default()
	router.Static("/public", "./public")
	//cors
	router.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowAllOrigins:  true,
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
	}))

	// rate limit
	router.Use(middleware.RateLimitMiddleware(time.Millisecond*10, 100))

	router.NoRoute(response.NotFound)
	//go-swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//Teapot
	router.GET("teapot", func(c *gin.Context) {
		c.JSON(http.StatusTeapot, gin.H{
			"message": "I'm a teapot",
			"story": "This code was defined in 1998 " +
				"as one of the traditional IETF April Fools' jokes," +
				" in RFC 2324, Hyper Text Coffee Pot Control Protocol," +
				" and is not expected to be implemented by actual HTTP servers." +
				" However, known implementations do exist.",
		})
	})
	router.GET("/", api.GetApiList)
	publicRouterV1 := router.Group("v1")
	{
		routers.InitBaseRouter(publicRouterV1)
		routers.InitVerificationCodeRouter(publicRouterV1)
	}

	privateGroupV1 := router.Group("v1")
	privateGroupV1.Use(middleware.JWTAuth())
	{
		routers.InitBookRouter(privateGroupV1)
		routers.InitFileRouter(privateGroupV1)
		routers.InitUserRouter(privateGroupV1)
	}

	return router
}
