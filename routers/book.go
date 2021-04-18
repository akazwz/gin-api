package routers

import (
	v1 "github.com/akaedison/go-gin-demo/api/v1"
	"github.com/gin-gonic/gin"
)

func InitBookRouter(Router *gin.RouterGroup) {
	BookRouter := Router.Group("books")
	{
		BookRouter.POST("", v1.CreateBook)
		BookRouter.DELETE("", v1.DeleteBook)
		BookRouter.PUT("", v1.UpdateBook)
		BookRouter.GET("", v1.GetBookList)
	}
}
