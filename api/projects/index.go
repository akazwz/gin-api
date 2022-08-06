package projects

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

func FindProjects(c *gin.Context) {
	// redis 中获取缓存
	result, err := global.GREDIS.Get(context.TODO(), "cache-projects").Result()
	// redis 中有缓存
	if err == nil {
		var projects []model.Project
		_ = json.Unmarshal([]byte(result), &projects)
		response.Ok(api.CodeCommonSuccess, projects, "success", c)
		return
	}
	projectService := service.ProjectService{}
	projects, err := projectService.FindProjects()
	if err != nil {
		response.BadRequest(api.CodeCommonFailed, nil, "获取失败", c)
		return
	}
	// posts 转为 json字符串
	bytes, _ := json.Marshal(projects)
	// 存入缓存
	_ = global.GREDIS.Set(context.TODO(), "cache-projects", string(bytes), 24*time.Hour).Err()

	response.Ok(api.CodeCommonSuccess, projects, "success", c)
}

func CreateProject(c *gin.Context) {
	uidAny, _ := c.Get("uid")
	uid := uidAny.(string)

	projectParams := request.CreateProjectsParams{}
	err := c.ShouldBind(&projectParams)
	if err != nil {
		response.BadRequest(api.CodeCommonFailed, nil, "参数错误", c)
		return
	}
	projectService := service.ProjectService{}
	project := model.Project{
		Name:    projectParams.Name,
		About:   projectParams.About,
		Website: projectParams.Website,
		Repo:    projectParams.Repo,
		Preview: projectParams.Preview,
		Readme:  projectParams.Readme,
		UID:     uid,
	}

	projectInstance, err := projectService.CreateProject(project)
	if err != nil {
		response.BadRequest(api.CodeCommonFailed, nil, "新建失败", c)
		return
	}
	// 清除 posts 缓存
	global.GREDIS.Del(context.TODO(), "cache-projects")
	response.Created(api.CodeCommonSuccess, projectInstance, "success", c)
}
