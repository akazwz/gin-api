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

func DeleteBook(c *gin.Context) {
	var reqId request.GetById
	if err := c.ShouldBindJSON(&reqId); err != nil {
		response.FailWithMessage("Bind Error", c)
		return
	}

	if reqId.Id == 0 {
		response.FailWithMessage("ID can not be null or 0", c)
		return
	}

	if err := service.DeleteBook(reqId.Id); err != nil {
		response.FailWithMessage("Delete Error", c)
	} else {
		response.OkWithMessage("Delete Success", c)
	}
}

func getUserId(c *gin.Context) uint {
	if claims, exists := c.Get("claims"); !exists {
		return 0
	} else {
		waitUse := claims.(*request.CustomClaims)
		return waitUse.ID
	}
}

func getUserUuid(c *gin.Context) string {
	if claims, exists := c.Get("claims"); !exists {
		return ""
	} else {
		waitUse := claims.(*request.CustomClaims)
		return waitUse.UUID.String()
	}
}
