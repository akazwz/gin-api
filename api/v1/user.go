package v1

import (
	"github.com/akaedison/go-gin-demo/model"
	"github.com/akaedison/go-gin-demo/model/request"
	"github.com/akaedison/go-gin-demo/model/response"
	"github.com/akaedison/go-gin-demo/service"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var login request.Login
	err := c.ShouldBindJSON(&login)
	if err != nil {
		response.FailWithMessage("Login Error", c)
		return
	}

	u := &model.User{Username: login.Username, Password: login.Password}
	if err, _ := service.Login(u); err != nil {
		response.FailWithMessage("Username Or Password Wrong", c)
		return
	}

	response.OkWithDetail(login, "Login Success", c)
}

func Register(c *gin.Context) {
	var register request.Register
	err := c.ShouldBindJSON(&register)
	if err != nil {
		response.FailWithMessage("Register Error", c)
		return
	}

	user := &model.User{
		Username:  register.Username,
		Password:  register.Password,
		NickName:  register.NickName,
		HeaderImg: register.HeaderImg,
	}
	err, _ = service.Register(*user)
	if err != nil {
		response.FailWithMessage("Register Failed", c)
		return
	}
	response.OkWithDetail(register, "Register Success", c)
}
