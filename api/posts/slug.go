package posts

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/akazwz/gin-api/model"
	"time"

	"github.com/akazwz/gin-api/api"
	"github.com/akazwz/gin-api/global"
	"github.com/akazwz/gin-api/model/response"
	"github.com/akazwz/gin-api/service"
	"github.com/gin-gonic/gin"
)

func GetPostById(c *gin.Context) {
	postService := service.PostService{}

	id := c.Param("id")

	cacheKey := fmt.Sprintf("%s-%s", "cache-post", id)
	// redis 中获取缓存
	result, err := global.GREDIS.Get(context.TODO(), cacheKey).Result()
	// redis 中有缓存
	if err == nil {
		var post model.Post
		_ = json.Unmarshal([]byte(result), &post)
		_ = postService.AddViewed(id)
		response.Ok(api.CodeCommonSuccess, post, "success", c)
		return
	}

	post, err := postService.FindPostByID(id)
	if err != nil {
		response.NotFound(api.CodeCommonFailed, "Not Found", c)
		return
	}
	_ = postService.AddViewed(id)

	// post 转为 json字符串
	bytes, _ := json.Marshal(post)
	// 存入缓存
	_ = global.GREDIS.Set(context.TODO(), cacheKey, string(bytes), 1*time.Minute).Err()

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
