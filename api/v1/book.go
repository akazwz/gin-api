package v1

import (
	"github.com/akaedison/go-gin-demo/model"
	"github.com/akaedison/go-gin-demo/model/request"
	"github.com/akaedison/go-gin-demo/model/response"
	"github.com/akaedison/go-gin-demo/service"
	"github.com/gin-gonic/gin"
)

func AddBook(c *gin.Context) {
	var book request.Book
	err := c.ShouldBindJSON(&book)
	if err != nil {
		response.FailWithMessage("Add Error err", c)
		return
	}

	b := &model.Book{
		BookName:     book.BookName,
		Author:       book.Author,
		Price:        book.Price,
		Introduction: book.Introduction,
	}
	if err, bookAdded := service.AddBook(b); err != nil {
		response.FailWithMessage("Add Error", c)
		return
	} else {
		response.OkWithDetail(bookAdded, "Add Success", c)
	}
}
