package request

type PageInfo struct {
	Page     int `json:"page" form:"page"`
	PageSize int `json:"page_size" form:"page_size"`
}

type GetById struct {
	Id float64 `json:"id" form:"id"`
}
