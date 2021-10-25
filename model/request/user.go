package request

import uuid "github.com/satori/go.uuid"

// Register 普通注册
type Register struct {
	Username         string `json:"username" form:"username" binding:"required"`
	Password         string `json:"password" form:"password" binding:"required"`
	NickName         string `json:"nick_name" form:"nick_name"`
	AvatarUrl        string `json:"avatar_url" form:"avatar_url"`
	Phone            string `json:"phone" form:"phone" binding:"required"`
	VerificationCode string `json:"verification_code" form:"verification_code" binding:"required"`
}

// LoginByUsernamePwd 用户名密码登录
type LoginByUsernamePwd struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

// LoginByPhonePwd 手机号密码登录
type LoginByPhonePwd struct {
	Phone    string `form:"phone" json:"phone" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

// LoginByPhoneVerificationCode 手机号验证码登录
type LoginByPhoneVerificationCode struct {
	Phone            string `form:"phone" json:"phone" binding:"required"`
	VerificationCode string `form:"verification_code" json:"verification_code" binding:"required"`
}

// LoginByOpenId 微信 小程序 openid 登录
type LoginByOpenId struct {
	Code    string `json:"code" form:"code"  binding:"required"`
	Encrypt string `json:"encrypt" form:"encrypt"`
	Iv      string `json:"iv" form:"iv"`
}

// ChangePassword 新旧密码普通修改密码
type ChangePassword struct {
	OldPassword string `json:"old_password" form:"old_password" binding:"required"`
	NewPassword string `json:"new_password" form:"new_password" binding:"required"`
}

// ChangePasswordByPhoneVerificationCode 手机号验证码修改密码
type ChangePasswordByPhoneVerificationCode struct {
	NewPassword      string `json:"new_password" form:"new_password" binding:"required"`
	Phone            string `json:"phone" form:"phone" binding:"required"`
	VerificationCode string `json:"verification_code" form:"verification_code" binding:"required"`
}

// UpdateUserProfile 修改账户信息
type UpdateUserProfile struct {
	NickName  string `json:"nick_name" form:"nick_name"`
	AvatarURL string `json:"avatar_url" form:"avatar_url"`
	Gender    int    `json:"gender" form:"avatar_url"`
	Bio       string `json:"bio" form:"bio"`
}

// SetUserAuth 设置用户权限
type SetUserAuth struct {
	UUID        uuid.UUID `json:"uuid" binding:"required"`
	AuthorityId string    `json:"authority_id" binding:"required"`
}
