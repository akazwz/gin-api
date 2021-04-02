package v1

import (
	"github.com/akaedison/go-gin-demo/model"
	"github.com/akaedison/go-gin-demo/model/request"
	"github.com/akaedison/go-gin-demo/model/response"
	"github.com/akaedison/go-gin-demo/service"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var Login request.Login
	_ = c.ShouldBindJSON(&Login)
	//username := c.PostForm("username")
	//password := c.PostForm("password")
	U := &model.User{Username: Login.Username, Password: Login.Password}
	if err, user := service.Login(U); err != nil {
		response.FailWithMessage("用户名不存在或者密码错误", c)
	} else {
		response.OkWithDetail(user, "登陆成功", c)
	}
}
