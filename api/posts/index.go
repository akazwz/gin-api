package posts

import (
	"context"
	"encoding/json"
	"time"

	"github.com/akazwz/gin-api/api"
	"github.com/akazwz/gin-api/global"
	"github.com/akazwz/gin-api/model"
	"github.com/akazwz/gin-api/model/request"
	"github.com/akazwz/gin-api/model/response"
	"github.com/akazwz/gin-api/service"
	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {
	uidAny, _ := c.Get("uid")
	uid := uidAny.(string)

	postParams := request.CreatePostParams{}
	err := c.ShouldBind(&postParams)
	if err != nil {
		response.BadRequest(api.CodeCommonFailed, nil, "参数错误", c)
		return
	}

	postService := service.PostService{}
	post := model.Post{
		Title:   postParams.Title,
		Cover:   postParams.Cover,
		Content: postParams.Content,
		Viewed:  0,
		UID:     uid,
	}
	postInstance, err := postService.CreatePost(post)
	if err != nil {
		response.BadRequest(api.CodeCommonFailed, nil, "新建失败", c)
		return
	}
	// 清除 posts 缓存
	global.GREDIS.Set(context.TODO(), "cache-posts", nil, 0)
	response.Created(api.CodeCommonSuccess, postInstance, "success", c)
}

func FindPosts(c *gin.Context) {
	// redis 中获取缓存
	result, err := global.GREDIS.Get(context.TODO(), "cache-posts").Result()
	// redis 中有缓存
	if err == nil {
		var posts []model.Post
		_ = json.Unmarshal([]byte(result), &posts)
		response.Ok(api.CodeCommonSuccess, posts, "success", c)
		return
	}

	postService := service.PostService{}
	posts, err := postService.FindPosts()
	if err != nil {
		response.BadRequest(api.CodeCommonFailed, nil, "获取失败", c)
		return
	}
	// posts 转为 json字符串
	bytes, _ := json.Marshal(posts)
	// 存入缓存
	_ = global.GREDIS.Set(context.TODO(), "cache-posts", string(bytes), 24*time.Hour).Err()

	response.Ok(api.CodeCommonSuccess, posts, "success", c)
}
