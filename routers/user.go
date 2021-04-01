package routers

import (
	"github.com/akaedison/go-gin-demo/global"
	"github.com/akaedison/go-gin-demo/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user")
	{
		user := &model.User{}
		UserRouter.GET("/get", func(c *gin.Context) {
			global.GDB.First(user)
			c.JSON(http.StatusOK, gin.H{
				"code": 200,
				"msg":  "success",
				"data": user,
			})
		})
	}
}
