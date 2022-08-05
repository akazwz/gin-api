package posts

import (
	"context"

	"github.com/akazwz/gin-api/api"
	"github.com/akazwz/gin-api/global"
	"github.com/akazwz/gin-api/model/response"
	"github.com/akazwz/gin-api/service"
	"github.com/gin-gonic/gin"
)

func GetPostById(c *gin.Context) {
	id := c.Param("id")
	postService := service.PostService{}
	post, err := postService.FindPostByID(id)
	if err != nil {
		response.NotFound(api.CodeCommonFailed, "Not Found", c)
		return
	}
	_ = postService.AddViewed(id)
	response.Ok(api.CodeCommonSuccess, post, "success", c)
}

func DeletePostById(c *gin.Context) {
	id := c.Param("id")
	postService := service.PostService{}
	err := postService.DeletePostByID(id)
	if err != nil {
		response.BadRequest(api.CodeCommonFailed, nil, "删除失败", c)
		return
	}
	// 清除 posts 缓存
	global.GREDIS.Set(context.TODO(), "cache-posts", nil, 0)
	response.Ok(api.CodeCommonSuccess, nil, "success", c)
}
