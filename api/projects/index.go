package projects

import (
	"time"

	"github.com/akazwz/gin-api/api"
	"github.com/akazwz/gin-api/model"
	"github.com/akazwz/gin-api/model/request"
	"github.com/akazwz/gin-api/model/response"
	"github.com/akazwz/gin-api/service"
	"github.com/akazwz/gin-api/utils"
	"github.com/gin-gonic/gin"
)

var projectService = service.ProjectService{}

func FindProjects(c *gin.Context) {
	// 获取缓存
	var projectsCache []model.Project
	err := utils.RedisCacheGet("cache-projects", &projectsCache)
	if err == nil {
		response.Ok(api.CodeCommonSuccess, projectsCache, "success", c)
		return
	}

	projects, err := projectService.FindProjects()
	if err != nil {
		response.BadRequest(api.CodeCommonFailed, nil, "获取失败", c)
		return
	}
	// 设置缓存
	_ = utils.RedisCacheSet("cache-projects", projects, 24*time.Hour)
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
	// 清除缓存
	_ = utils.RedisCacheDel("cache-projects")
	response.Created(api.CodeCommonSuccess, projectInstance, "success", c)
}
