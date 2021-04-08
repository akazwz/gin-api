package request

type Register struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	NickName  string `json:"nick_name"`
	HeaderImg string `json:"header_img"`
}

type Login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"-"`
}
