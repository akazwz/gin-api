package request

import uuid "github.com/satori/go.uuid"

type Register struct {
	Username         string `json:"username" form:"username" binding:"required"`
	Password         string `json:"password" form:"password" binding:"required"`
	NickName         string `json:"nick_name" form:"nick_name"`
	HeaderImg        string `json:"header_img" form:"header"`
	Phone            string `json:"phone" form:"phone" binding:"required"`
	VerificationCode string `json:"verification_code" form:"verification_code" binding:"required"`
}

type LoginByUsernamePwd struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type LoginByPhonePwd struct {
	Phone    string `form:"phone" json:"phone" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type LoginByPhoneVerificationCode struct {
	Phone            string `form:"phone" json:"phone" binding:"required"`
	VerificationCode string `form:"verification_code" json:"verification_code" binding:"required"`
}

type ChangePassword struct {
	Username    string `json:"username" form:"username" binding:"required"`
	OldPassword string `json:"old_password" form:"old_password" binding:"required"`
	NewPassword string `json:"new_password" form:"new_password" binding:"required"`
}

type SetUserAuth struct {
	UUID        uuid.UUID `json:"uuid" binding:"required"`
	AuthorityId string    `json:"authority_id" binding:"required"`
}
