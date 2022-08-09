package projects

import (
	"fmt"
	"time"

	"github.com/akazwz/gin-api/api"
	"github.com/akazwz/gin-api/model"
	"github.com/akazwz/gin-api/model/response"
	"github.com/akazwz/gin-api/utils"
	"github.com/gin-gonic/gin"
)

func FindProjectByID(c *gin.Context) {
	id := c.Param("id")
	// 缓存中取
	key := fmt.Sprintf("cache-project-%s", id)
	cacheProject := model.Project{}
	err := utils.RedisCacheGet(key, &cacheProject)
	if err == nil {
		response.Ok(api.CodeCommonSuccess, cacheProject, "success-cache", c)
		return
	}
	project, err := projectService.FindProjectByID(id)
	if err != nil {
		response.NotFound(api.CodeCommonFailed, "Not Found", c)
		return
	}
	// 放入缓存
	_ = utils.RedisCacheSet(key, project, 24*time.Hour)
	response.Ok(api.CodeCommonSuccess, project, "success", c)
}

func DeleteProjectByID(c *gin.Context) {
	id := c.Param("id")
	key := fmt.Sprintf("cache-project-%s", id)
	err := projectService.DeleteProjectByID(id)
	if err != nil {
		response.BadRequest(api.CodeCommonFailed, nil, "删除失败", c)
		return
	}
	// 清除缓存
	_ = utils.RedisCacheDel("cache-projects")
	_ = utils.RedisCacheDel(key)

	response.Ok(api.CodeCommonSuccess, nil, "success", c)
	return
}
