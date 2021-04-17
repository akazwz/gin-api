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
