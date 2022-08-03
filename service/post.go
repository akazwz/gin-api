package service

import (
	"github.com/akazwz/gin-api/global"
	"github.com/akazwz/gin-api/model"
	uuid "github.com/satori/go.uuid"
)

type PostService struct{}

func (postService *PostService) CreatePost(post model.Post) (model.Post, error) {
	post.UUID = uuid.NewV4().String()
	err := global.GDB.Create(&post).Error
	return post, err
}

func (postService *PostService) FindPosts() ([]model.Post, error) {
	var posts []model.Post
	err := global.GDB.Find(&posts).Error
	return posts, err
}

func (postService *PostService) FindPostByID(id string) (model.Post, error) {
	var post model.Post
	err := global.GDB.Where("uuid = ?", id).First(&post).Error
	return post, err
}

func (postService *PostService) DeletePostByID(id string) error {
	var post model.Post
	err := global.GDB.Where("uuid = ?", id).Delete(&post).Error
	return err
}

func (postService *PostService) AddViewed(id string) error {
	var post model.Post
	err := global.GDB.Where("uuid = ?", id).First(&post).Update("viewed", post.Viewed+1).Error
	return err
}
