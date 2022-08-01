package auth

import (
	"github.com/akazwz/gin-api/api"
	"github.com/akazwz/gin-api/model/response"
	"github.com/akazwz/gin-api/service"
	"github.com/gin-gonic/gin"
)

func Me(c *gin.Context) {
	authService := service.AuthService{}

	uidAny, _ := c.Get("uid")
	uid := uidAny.(string)

	user := authService.FindUserByUID(uid)
	/* 用户不存在 */
	if user == nil {
		response.BadRequest(api.CodeCommonFailed, nil, "账户不存在", c)
		return
	}
	response.Ok(api.CodeCommonSuccess, gin.H{
		"user": user,
	}, "success", c)
}
