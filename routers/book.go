package routers

import (
	v1 "github.com/akaedison/go-gin-demo/api/v1"
	"github.com/gin-gonic/gin"
)

func InitBookRouter(Router *gin.RouterGroup) {
	BookRouter := Router.Group("book")
	{
		BookRouter.POST("/add", v1.AddBook)
	}
}