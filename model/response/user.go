package response

// UserResponse 返回的user, 去除敏感字段
type UserResponse struct {
	Username    string `json:"username"`
	NickName    string `json:"nick_name"`
	AvatarUrl   string `json:"avatar_url"`
	Phone       string `json:"phone"`
	Gender      int    `json:"gender"`
	Bio         string `json:"bio"`
	AuthorityId string `json:"authority_id"`
}

// LoginResponse 登录返回,user token 和过期时间
type LoginResponse struct {
	User      UserResponse `json:"user"`
	Token     string       `json:"token"`
	ExpiresAt int64        `json:"expires_at"`
}
