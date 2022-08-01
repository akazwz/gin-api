package service

import (
	"errors"

	"github.com/akazwz/gin-api/global"
	"github.com/akazwz/gin-api/model"
	"github.com/akazwz/gin-api/utils"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type AuthService struct{}

func (authService *AuthService) SignupService(u model.User) (*model.User, error) {
	user := authService.FindUserByUsername(u.Username)
	// username 已经存在
	if user != nil {
		return user, errors.New("用户名已注册")
	}
	u.Password = utils.BcryptHash(u.Password)
	u.UUID = uuid.NewV4()
	err := global.GDB.Create(&u).Error
	return &u, err
}

func (authService *AuthService) LoginService(u model.User) (*model.User, error) {
	user := authService.FindUserByUsername(u.Username)
	if user == nil {
		return nil, errors.New("用户不存在")
	}
	isPasswordCorrect := utils.BcryptCheck(u.Password, user.Password)
	if !isPasswordCorrect {
		return nil, errors.New("密码错误")
	}
	return user, nil
}

func (authService *AuthService) FindUserByUsername(username string) *model.User {
	var user model.User
	err := global.GDB.Where("username = ?", username).First(&user).Error
	// 用户不存在
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}
	return &user
}

func (authService *AuthService) FindUserByUID(uid string) *model.User {
	var user model.User
	err := global.GDB.Where("uuid = ?", uid).First(&user).Error
	// 用户不存在
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}
	return &user
}
