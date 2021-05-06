package request

type Book struct {
	BookName     string  `json:"book_name" form:"book_name" binding:"required"`
	Author       string  `json:"author" form:"book_name"`
	Price        float64 `json:"price" form:"price"`
	Cover        string  `json:"cover" form:"cover"`
	Introduction string  `json:"introduction" form:"introduction"`
}
