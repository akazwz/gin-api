package v1

import (
	"time"

	"github.com/akazwz/go-gin-restful-api/global"
	"github.com/akazwz/go-gin-restful-api/middleware"
	"github.com/akazwz/go-gin-restful-api/model"
	"github.com/akazwz/go-gin-restful-api/model/request"
	"github.com/akazwz/go-gin-restful-api/model/response"
	"github.com/akazwz/go-gin-restful-api/pkg/utils"
	"github.com/akazwz/go-gin-restful-api/service"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// CreateTokenByUsernamePwd
// @Summary Create A Token By Username Pwd
// @Title Create Token
// @Author zwz
// @Description create token
// @Tags token
// @Accept json
// @Produce json
// @Param loginByUp body request.LoginByUsernamePwd true "login by username pwd"
// @Success 201 {object} response.LoginResponse
// @Failure 400 {object} response.Response
// @Router /token/username-pwd [post]
func CreateTokenByUsernamePwd(c *gin.Context) {
	var login request.LoginByUsernamePwd

	if err := c.ShouldBindJSON(&login); err != nil {
		response.CommonFailed("Bind Json Error", CodeBindError, c)
		return
	}

	u := &model.User{Username: login.Username, Password: login.Password}
	if err, user := service.LoginByUsernamePwd(u); err != nil {
		response.CommonFailed("Username Or Password Error", CodeDbErr, c)
		return
	} else {
		TokenNext(c, *user)
	}
}

// CreateTokenByPhonePwd
// @Summary Create A Token By Phone Pwd
// @Title Create Token
// @Author zwz
// @Description create token
// @Tags token
// @Accept json
// @Produce json
// @Param loginByUp body request.LoginByPhonePwd true "login by phone pwd"
// @Success 201 {object} response.LoginResponse
// @Failure 400 {object} response.Response
// @Router /token/phone-pwd [post]
func CreateTokenByPhonePwd(c *gin.Context) {
	var login request.LoginByPhonePwd

	if err := c.ShouldBindJSON(&login); err != nil {
		response.CommonFailed("Bind Json Error", CodeBindError, c)
		return
	}

	// 判断手机号是否存在
	exist, user := service.IsPhoneExist(login.Phone)
	if !exist {
		response.CommonFailed("No Such Phone", CodeNoSuchPhoneError, c)
		return
	}

	// 判断密码是否正确
	if user.Password != login.Password {
		response.CommonFailed("Phone Or Password Error", CodeDbErr, c)
		return
	} else {
		TokenNext(c, *user)
	}
}

// CreateTokenByPhoneVerificationCode
// @Summary Create A Token By Phone Code
// @Title Create Token
// @Author zwz
// @Description create token by phone code
// @Tags token
// @Accept json
// @Produce json
// @Param loginByPc body request.LoginByPhoneVerificationCode true "login by phone verification code"
// @Success 201 {object} response.LoginResponse
// @Failure 400 {object} response.Response
// @Router /token/phone-code [post]
func CreateTokenByPhoneVerificationCode(c *gin.Context) {
	var login request.LoginByPhoneVerificationCode

	if err := c.ShouldBindJSON(&login); err != nil {
		response.CommonFailed("Bind Json Error", CodeBindError, c)
		return
	}

	// 判断手机号是否存在
	exist, user := service.IsPhoneExist(login.Phone)
	if !exist {
		response.CommonFailed("No Such Phone", CodeNoSuchPhoneError, c)
		return
	}

	passed := utils.GetVerificationStatus(login.Phone, login.VerificationCode, c.Request.Context())
	if !passed {
		response.CommonFailed("verification code error", CodeVerificationCodeError, c)
		return
	} else {
		TokenNext(c, *user)
	}
}

func CreateTokenByOpenId(c *gin.Context) {
	var login request.LoginByOpenId

	if err := c.ShouldBindJSON(&login); err != nil {
		response.CommonFailed("Bind Json Error", CodeBindError, c)
		return
	}

	session := utils.GetSessionByCode(login.Code)

	// 判断OpenId是否存在
	exist, user := service.IsOpenIdExist(session.OpenID)
	if !exist {
		// 不存在新建用户
		response.CommonFailed("No Such Phone", CodeNoSuchPhoneError, c)
		return
	} else {
		// 存在返回token
		TokenNext(c, *user)
	}
}

// TokenNext
// generate and return token
func TokenNext(c *gin.Context, user model.User) {
	j := &middleware.JWT{SigningKey: []byte(global.CFG.JWT.SigningKey)}
	claims := request.CustomClaims{
		UUID:       user.UUID,
		ID:         user.ID,
		Username:   user.Username,
		Phone:      user.Phone,
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

	u := model.User{
		Username:    user.Username,
		AvatarUrl:   user.AvatarUrl,
		NickName:    user.NickName,
		AuthorityId: user.AuthorityId,
	}

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

	// check phone exists
	existPhone, _ := service.IsPhoneExist(register.Phone)
	if existPhone {
		response.CommonFailed("Phone Already Exits", CodePhoneAlreadyExitsError, c)
		return
	}

	// verification code check
	passed := utils.GetVerificationStatus(register.Phone, register.VerificationCode, c.Request.Context())
	if !passed {
		response.CommonFailed("verification code error", CodeVerificationCodeError, c)
		return
	}

	user := &model.User{
		Username:  register.Username,
		Phone:     register.Phone,
		Password:  register.Password,
		NickName:  register.NickName,
		AvatarUrl: register.AvatarUrl,
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
// @Router /users/password [patch]
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
		response.SuccessWithMessage("Password Change Success", c)
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

// SetUserAuthority
// @Summary Set UserAuthority
// @Title Set UserAuthority
// @Author zwz
// @Description set userAuthority
// @Tags user
// @Accept json
// @Produce json
// @Param setUserAuth body request.SetUserAuth true "setUserAuth"
// @Param token header string true "token"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /users/authority [patch]
func SetUserAuthority(c *gin.Context) {
	var sua request.SetUserAuth
	if err := c.ShouldBindJSON(&sua); err != nil {
		response.CommonFailed("Bind Json Error", CodeBindError, c)
		return
	}

	if err := service.SetUserAuthority(sua.UUID, sua.AuthorityId); err != nil {
		response.CommonFailed("Set User Authority Error", CodeDbErr, c)
	} else {
		response.SuccessWithMessage("Set User Authority Success", c)
	}
}
