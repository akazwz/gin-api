package posts

import (
	"fmt"
	"time"

	"github.com/akazwz/gin-api/api"
	"github.com/akazwz/gin-api/model"
	"github.com/akazwz/gin-api/model/response"
	"github.com/akazwz/gin-api/service"
	"github.com/akazwz/gin-api/utils"
	"github.com/gin-gonic/gin"
)

func GetPostById(c *gin.Context) {
	id := c.Param("id")
	key := fmt.Sprintf("%s-%s", "cache-post-", id)

	// 取缓存
	cachePost := model.Post{}
	err := utils.RedisCacheGet(key, &cachePost)
	if err == nil {
		response.Ok(api.CodeCommonSuccess, cachePost, "success-cache", c)
		return
	}

	postService := service.PostService{}
	post, err := postService.FindPostByID(id)
	if err != nil {
		response.NotFound(api.CodeCommonFailed, "Not Found", c)
		return
	}
	// 放缓存
	_ = utils.RedisCacheSet(key, post, 24*time.Hour)
	response.Ok(api.CodeCommonSuccess, post, "success", c)
}

func DeletePostById(c *gin.Context) {
	id := c.Param("id")
	key := fmt.Sprintf("%s-%s", "cache-post-", id)

	postService := service.PostService{}
	err := postService.DeletePostByID(id)
	if err != nil {
		response.BadRequest(api.CodeCommonFailed, nil, "删除失败", c)
		return
	}
	// 清除 posts 缓存
	_ = utils.RedisCacheDel("cache-posts")
	_ = utils.RedisCacheDel(key)
	response.Ok(api.CodeCommonSuccess, nil, "success", c)
}
