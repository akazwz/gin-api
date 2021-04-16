package v1

import (
	"github.com/akaedison/go-gin-demo/global"
	"github.com/akaedison/go-gin-demo/middleware"
	"github.com/akaedison/go-gin-demo/model"
	"github.com/akaedison/go-gin-demo/model/request"
	"github.com/akaedison/go-gin-demo/model/response"
	"github.com/akaedison/go-gin-demo/service"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"time"
)

func Login(c *gin.Context) {
	var login request.Login
	err := c.ShouldBindJSON(&login)
	if err != nil {
		response.FailWithMessage("Login Error", c)
		return
	}

	u := &model.User{Username: login.Username, Password: login.Password}
	if err, user := service.Login(u); err != nil {
		response.FailWithMessage("Username Or Password Wrong", c)
		return
	} else {
		TokenNext(c, *user)
	}

}

func TokenNext(c *gin.Context, user model.User) {
	j := &middleware.JWT{SigningKey: []byte(global.CFG.JWT.SigningKey)}
	claims := request.CustomClaims{
		UUID:       uuid.UUID{},
		ID:         user.ID,
		Username:   user.Username,
		NickName:   user.NickName,
		BufferTime: global.CFG.JWT.BufferTime,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,
			ExpiresAt: time.Now().Unix() + global.CFG.JWT.ExpiresTime,
			Issuer:    "zwz",
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		response.FailWithMessage("get token error", c)
		return
	}

	response.OkWithDetail(response.LoginResponse{
		User:      user,
		Token:     token,
		ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
	}, "Login Success", c)
}

func Register(c *gin.Context) {
	var register request.Register
	err := c.ShouldBindJSON(&register)
	if err != nil {
		response.FailWithMessage("Register Error", c)
		return
	}

	user := &model.User{
		Username:  register.Username,
		Password:  register.Password,
		NickName:  register.NickName,
		HeaderImg: register.HeaderImg,
	}
	err, _ = service.Register(*user)
	if err != nil {
		response.FailWithMessage("Register Failed", c)
		return
	}
	response.OkWithDetail(register, "Register Success", c)
}
