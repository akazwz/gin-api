package projects

import (
	"context"
	"github.com/akazwz/gin-api/api"
	"github.com/akazwz/gin-api/global"
	"github.com/akazwz/gin-api/model/response"
	"github.com/gin-gonic/gin"
)

func FindProjectByID(c *gin.Context) {
	id := c.Param("id")
	project, err := projectService.FindProjectByID(id)
	if err != nil {
		response.NotFound(api.CodeCommonFailed, "Not Found", c)
		return
	}
	response.Ok(api.CodeCommonSuccess, project, "success", c)
}

func DeleteProjectByID(c *gin.Context) {
	id := c.Param("id")
	err := projectService.DeleteProjectByID(id)
	if err != nil {
		response.BadRequest(api.CodeCommonFailed, nil, "删除失败", c)
		return
	}
	// 清除 projects 缓存
	global.GREDIS.Del(context.TODO(), "cache-projects")
	response.Ok(api.CodeCommonSuccess, nil, "success", c)
	return
}
