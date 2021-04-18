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
	"time"
)

// CreateToken
// @Summary Create A Token
// @Title Create Token
// @Author zwz
// @Description create book
// @Tags token
// @Accept json
// @Produce json
// @Param login body request.Login true "login"
// @Success 201 {object} response.LoginResponse
// @Failure 400 {object} response.Response
// @Router /token [post]
func CreateToken(c *gin.Context) {
	var login request.Login

	if err := c.ShouldBindJSON(&login); err != nil {
		response.CommonFailed("Bind Json Error", CodeBindError, c)
		return
	}

	u := &model.User{Username: login.Username, Password: login.Password}
	if err, user := service.Login(u); err != nil {
		response.CommonFailed("Username Or Password Error", CodeDbErr, c)
		return
	} else {
		TokenNext(c, *user)
	}

}

func TokenNext(c *gin.Context, user model.User) {
	j := &middleware.JWT{SigningKey: []byte(global.CFG.JWT.SigningKey)}
	claims := request.CustomClaims{
		UUID:       user.UUID,
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
		response.CommonFailed("Create Token Error", CodeCreateTokenError, c)
		return
	}

	response.Created(response.LoginResponse{
		User:      user,
		Token:     token,
		ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
	}, "Login Success", c)
}

// CreateUser
// @Summary Create A User
// @Title Create User
// @Author zwz
// @Description create user
// @Tags user
// @Accept json
// @Produce json
// @Param user body request.Register true "user"
// @Success 201 {object} model.Book
// @Failure 400 {object} response.Response
// @Router /users [post]
func CreateUser(c *gin.Context) {
	var register request.Register
	err := c.ShouldBindJSON(&register)
	if err != nil {
		response.CommonFailed("Bind Json Error", CodeBindError, c)
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
		response.CommonFailed("Register Failed", CodeDbErr, c)
		return
	}
	response.Created(register, "Register Success", c)
}
