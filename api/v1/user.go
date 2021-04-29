package v1

import (
	"github.com/akazwz/go-gin-demo/global"
	"github.com/akazwz/go-gin-demo/middleware"
	"github.com/akazwz/go-gin-demo/model"
	"github.com/akazwz/go-gin-demo/model/request"
	"github.com/akazwz/go-gin-demo/model/response"
	"github.com/akazwz/go-gin-demo/service"
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

	u := model.User{Username: user.Username}

	response.Created(response.LoginResponse{
		User:      u,
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

// ChangePassword
// @Summary Change Password
// @Title Change Password
// @Author zwz
// @Description change password
// @Tags user
// @Accept json
// @Produce json
// @Param changePassword body request.ChangePassword true "ChangePassword"
// @Param token header string true "token"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /users/password [put]
func ChangePassword(c *gin.Context) {
	var changePassword request.ChangePassword
	if err := c.ShouldBindJSON(&changePassword); err != nil {
		response.CommonFailed("Bind Json Error", CodeBindError, c)
	}
	u := model.User{Username: changePassword.Username, Password: changePassword.OldPassword}
	if err, _ := service.ChangePassword(&u, changePassword.NewPassword); err != nil {
		response.CommonFailed("Change Password Error, Password Not Correct", CodeDbErr, c)
	} else {
		u = model.User{Username: changePassword.Username, Password: changePassword.NewPassword}
		response.CommonSuccess(0, u, "Password Change Success", c)
	}
}

// GetUserList
// @Summary Get UserList
// @Title Get UserList
// @Author zwz
// @Description get user list
// @Tags user
// @Accept json
// @Produce json
// @Param page query int true "page"
// @Param page_size query int true "page_size"
// @Param token header string true "token"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /users [get]
func GetUserList(c *gin.Context) {
	var pageInfo request.PageInfo
	if err := c.ShouldBindQuery(&pageInfo); err != nil {
		response.CommonFailed("Bind Query Error", CodeBindError, c)
		return
	}

	if err, list, total := service.GetUserInfoList(pageInfo); err != nil {
		response.CommonFailed("Get UserList Error", CodeDbErr, c)
	} else {
		response.CommonSuccess(0, response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "Get UserList Success", c)
	}
}
