package posts

import (
	"github.com/akazwz/gin-api/api"
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
	response.Created(api.CodeCommonSuccess, postInstance, "success", c)
}

func FindPosts(c *gin.Context) {
	postService := service.PostService{}
	posts, err := postService.FindPosts()
	if err != nil {
		response.BadRequest(api.CodeCommonFailed, nil, "获取失败", c)
		return
	}
	response.Ok(api.CodeCommonSuccess, posts, "success", c)
}
