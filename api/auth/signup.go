package auth

import (
	"github.com/akazwz/gin-api/api"
	"github.com/akazwz/gin-api/model"
	"github.com/akazwz/gin-api/model/request"
	"github.com/akazwz/gin-api/model/response"
	"github.com/akazwz/gin-api/service"
	"github.com/gin-gonic/gin"
)

// SignupByUsernamePwd 注册(用户名 + 密码)
func SignupByUsernamePwd(c *gin.Context) {
	authService := service.AuthService{}
	// 获取绑定参数
	signup := request.SignUp{}
	err := c.ShouldBind(&signup)
	if err != nil {
		response.BadRequest(api.CodeCommonFailed, nil, "参数错误", c)
		return
	}
	// 构造用户
	user := model.User{
		Username: signup.Username,
		Password: signup.Password,
	}
	// 注册
	userInstance, err := authService.SignupService(user)
	if err != nil {
		response.BadRequest(api.CodeCommonFailed, nil, err.Error(), c)
		return
	}
	response.Created(api.CodeCommonSuccess, userInstance, "注册成功", c)
}
