package v1

import (
	"log"
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
// @Summary 用户名密码登录
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
// @Summary 手机号密码登录
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
// @Summary 手机号验证码登录
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

	// 检查验证码是否正确
	passed := utils.GetVerificationStatus(login.Phone, login.VerificationCode, c.Request.Context())
	if !passed {
		response.CommonFailed("verification code error", CodeVerificationCodeError, c)
		return
	} else {
		// 判断手机号是否存在
		exist, user := service.IsPhoneExist(login.Phone)
		if !exist {
			// 不存在新建用户
			userNew := model.User{
				Username: login.Phone,
				Phone:    login.Phone,
			}
			err, userRegister := service.RegisterByPhoneVerificationCode(userNew)
			if err != nil {
				response.CommonFailed("Register Error", CodeDbErr, c)
				return
			}
			// 新建成功,返回token
			TokenNext(c, *userRegister)
			return
		}
		// 存在返回token
		TokenNext(c, *user)
	}
}

// CreateTokenByOpenId
// @Summary 小程序 openid 登录
// @Title Create Token
// @Author zwz
// @Description create token by open id
// @Tags token
// @Accept json
// @Produce json
// @Param code body request.LoginByOpenId true "code"
// @Success 201 {object} response.LoginResponse
// @Failure 400 {object} response.Response
// @Router /token/open-id [post]
func CreateTokenByOpenId(c *gin.Context) {
	var login request.LoginByOpenId
	if err := c.ShouldBindJSON(&login); err != nil {
		response.CommonFailed("Bind Json Error", CodeBindError, c)
		return
	}
	session, err := utils.GetSessionByCode(login.Code)
	if err != nil {
		response.CommonFailed("Generate Session Error", CodeGenerateSessionError, c)
		return
	}
	// 判断OpenId是否存在
	exist, user := service.IsOpenIdExist(session.OpenID)
	// 不存在新建用户
	if !exist {
		// 获取 userInfo
		userInfo, err := utils.GetMiniUserInfo(session.SessionKey, login.Encrypt, login.Iv)
		if err != nil {
			log.Println(err)
			log.Println("get mini userinfo error")
			response.CommonFailed("Get Mini Userinfo error", CodeGetMiniUserInfoError, c)
			return
		}
		log.Println("phone" + userInfo.PhoneNumber)
		userNew := &model.User{
			Username:  session.OpenID,
			NickName:  userInfo.NickName,
			AvatarUrl: userInfo.AvatarURL,
			OpenId:    session.OpenID,
		}
		err, userRegister := service.RegisterByOpenId(*userNew)
		if err != nil {
			log.Println(err)
			response.CommonFailed("Register Failed", CodeDbErr, c)
			return
		}
		// 返回token
		TokenNext(c, *userRegister)
		return
	} else {
		// 存在直接返回token
		TokenNext(c, *user)
	}
}

// TokenNext 生成返回token
// generate and return token
func TokenNext(c *gin.Context, user model.User) {
	j := &middleware.JWT{SigningKey: []byte(global.CFG.JWT.SigningKey)}
	claims := request.CustomClaims{
		UUID:       user.UUID,
		ID:         user.ID,
		Username:   user.Username,
		Phone:      user.Phone,
		OpenId:     user.OpenId,
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

	userResponse := response.UserResponse{
		Username:    user.Username,
		AvatarUrl:   user.AvatarUrl,
		NickName:    user.NickName,
		Phone:       user.Phone,
		Gender:      user.Gender,
		Bio:         user.Bio,
		AuthorityId: user.AuthorityId,
	}

	response.Created(response.LoginResponse{
		User:      userResponse,
		Token:     token,
		ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
	}, "Login Success", c)
}

// CreateUser
// @Summary 普通注册
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

	// 检查手机号是否存在
	existPhone, _ := service.IsPhoneExist(register.Phone)
	if existPhone {
		// 手机号已经存在,返回
		response.CommonFailed("Phone Already Exits", CodePhoneAlreadyExitsError, c)
		return
	}

	// 检查验证码
	passed := utils.GetVerificationStatus(register.Phone, register.VerificationCode, c.Request.Context())
	if !passed {
		// 验证码错误
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
	// 注册
	err, _ = service.Register(*user)
	if err != nil {
		response.CommonFailed("Register Failed", CodeDbErr, c)
		return
	}
	userResponse := response.UserResponse{
		Username:  register.Username,
		NickName:  register.NickName,
		AvatarUrl: register.AvatarUrl,
	}
	response.Created(userResponse, "Register Success", c)
}

// ChangePassword
// @Summary 新旧密码普通修改密码
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
	// 获取 user uuid
	claims, _ := c.Get("claims")
	// convent claims to type *request.CustomClaims
	customClaims := claims.(*request.CustomClaims)
	// get user uuid to store who upload this file
	userUUID := customClaims.UUID
	u := model.User{UUID: userUUID, Password: changePassword.OldPassword}
	// 通过 user uuid 新旧密码修改密码
	if err, _ := service.ChangePassword(&u, changePassword.NewPassword); err != nil {
		response.CommonFailed("Change Password Error, Password Not Correct", CodeDbErr, c)
	} else {
		response.SuccessWithMessage("Password Change Success", c)
	}
}

// ChangePasswordByPhoneVerificationCode
// @Summary 手机号验证码修改密码
// @Title Change Password
// @Author zwz
// @Description change password
// @Tags user
// @Accept json
// @Produce json
// @Param changePassword body request.ChangePasswordByPhoneVerificationCode true "ChangePassword"
// @Param token header string true "token"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /users/password/phone-code [put]
func ChangePasswordByPhoneVerificationCode(c *gin.Context) {
	var changePassword request.ChangePasswordByPhoneVerificationCode
	if err := c.ShouldBindJSON(&changePassword); err != nil {
		response.CommonFailed("Bind Json Error", CodeBindError, c)
	}
	// 检查验证码
	passed := utils.GetVerificationStatus(changePassword.Phone, changePassword.VerificationCode, c.Request.Context())
	if !passed {
		// 验证码错误
		response.CommonFailed("verification code error", CodeVerificationCodeError, c)
		return
	}

	// 获取 user uuid
	claims, _ := c.Get("claims")
	// convent claims to type *request.CustomClaims
	customClaims := claims.(*request.CustomClaims)
	// get user uuid to store who upload this file
	userUUID := customClaims.UUID
	u := model.User{UUID: userUUID}

	// 通过 uuid 直接修改密码
	if err, _ := service.ChangePasswordByPhoneVerificationCode(&u, changePassword.NewPassword); err != nil {
		response.CommonFailed("Change Password Error, Password Not Correct", CodeDbErr, c)
	} else {
		response.SuccessWithMessage("Password Change Success", c)
	}
}

// UpdateUserProfile 修改账户资料
// @Summary 修改账号资料
// @Title 修改账户资料
// @Author zwz
// @Description change password
// @Tags user
// @Accept json
// @Produce json
// @Param 修改账户资料 body request.UpdateUserProfile true "修改账户资料"
// @Param token header string true "token"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /users/profile [put]
func UpdateUserProfile(c *gin.Context) {
	var profile request.UpdateUserProfile
	if err := c.ShouldBindJSON(&profile); err != nil {
		response.CommonFailed("Bind Json Error", CodeBindError, c)
	}
	// 获取 user uuid
	claims, _ := c.Get("claims")
	// convent claims to type *request.CustomClaims
	customClaims := claims.(*request.CustomClaims)
	// get user uuid to store who upload this file
	userUUID := customClaims.UUID
	userUpdate := model.User{
		UUID:      userUUID,
		NickName:  profile.NickName,
		AvatarUrl: profile.AvatarURL,
		Gender:    profile.Gender,
		Bio:       profile.Bio,
	}
	// 修改账户资料
	if err, _ := service.UpdateUserProfileByUser(&userUpdate); err != nil {
		response.CommonFailed("Update UserProfile Error", CodeDbErr, c)
	} else {
		response.SuccessWithMessage("Update User Profile Success", c)
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
// @Router /users/authority [put]
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
