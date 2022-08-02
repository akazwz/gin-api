package request

type CreatePostParams struct {
	Title   string `json:"title" form:"title" binding:"required"`
	Cover   string `json:"cover" form:"cover" binding:"required"`
	Content string `json:"content" form:"content" binding:"required"`
}
