package service

import (
	"github.com/akazwz/gin-api/global"
	"github.com/akazwz/gin-api/model"
	uuid "github.com/satori/go.uuid"
)

type ProjectService struct{}

func (p *ProjectService) CreateProject(project model.Project) (model.Project, error) {
	project.UUID = uuid.NewV4().String()
	err := global.GDB.Create(&project).Error
	return project, err
}

func (p *ProjectService) FindProjects() ([]model.Project, error) {
	var projects []model.Project
	err := global.GDB.Find(&projects).Error
	return projects, err
}

func (p *ProjectService) FindProjectByID(id string) (model.Project, error) {
	var project model.Project
	err := global.GDB.Where("uuid = ?", id).First(&project).Error
	return project, err
}

func (p *ProjectService) DeleteProjectByID(id string) error {
	var project model.Project
	err := global.GDB.Where("uuid = ?", id).Delete(&project).Error
	return err
}
