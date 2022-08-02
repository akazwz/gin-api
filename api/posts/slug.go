package posts

import (
	"github.com/akazwz/gin-api/api"
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
