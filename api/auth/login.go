package auth

import (
	"github.com/akazwz/gin-api/api"
	"github.com/akazwz/gin-api/model"
	"github.com/akazwz/gin-api/model/request"
	"github.com/akazwz/gin-api/model/response"
	"github.com/akazwz/gin-api/service"
	"github.com/akazwz/gin-api/utils"
	"github.com/gin-gonic/gin"
)

// LoginByUsernamePwd 登录(用户名 + 密码)
func LoginByUsernamePwd(c *gin.Context) {
	authService := service.AuthService{}
	// 获取绑定参数
	login := request.Login{}
	err := c.ShouldBind(&login)
	if err != nil {
		response.BadRequest(api.CodeCommonFailed, nil, "参数错误", c)
		return
	}
	// 构造用户
	user := model.User{
		Username: login.Username,
		Password: login.Password,
	}
	// 登录
	userInstance, err := authService.LoginService(user)
	if err != nil {
		response.BadRequest(api.CodeCommonFailed, nil, err.Error(), c)
		return
	}
	TokenNext(*userInstance, c)
}

func TokenNext(user model.User, c *gin.Context) {
	j := utils.NewJWT()
	claims := j.CreateClaims(request.BaseClaims{UUID: user.UUID})
	token, err := j.CreateToken(claims)
	if err != nil {
		response.BadRequest(api.CodeCommonFailed, nil, "获取token失败", c)
		return
	}
	response.Created(api.CodeCommonSuccess, gin.H{
		"user":       user,
		"token":      token,
		"expires_at": claims.RegisteredClaims.ExpiresAt,
	}, "登录成功", c)
}
