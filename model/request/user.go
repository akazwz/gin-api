package request

type Register struct {
	Username  string `json:"username" form:"username" binding:"required"`
	Password  string `json:"password" form:"password" binding:"required"`
	NickName  string `json:"nick_name" form:"nick_name"`
	HeaderImg string `json:"header_img" form:"header"`
}

type Login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type ChangePassword struct {
	Username    string `json:"username" form:"username" binding:"required"`
	OldPassword string `json:"old_password" form:"old_password" binding:"required"`
	NewPassword string `json:"new_password" form:"new_password" binding:"required"`
}
