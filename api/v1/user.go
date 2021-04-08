package v1

import (
	"github.com/akaedison/go-gin-demo/model/request"
	"github.com/akaedison/go-gin-demo/model/response"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var login request.Login
	err := c.ShouldBindJSON(&login)
	if err != nil {
		response.FailWithMessage("Bind Error", c)
		return
	}
	response.OkWithDetail(login, "Login Success", c)
}
