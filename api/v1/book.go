package v1

import (
	"github.com/akaedison/go-gin-demo/model"
	"github.com/akaedison/go-gin-demo/model/request"
	"github.com/akaedison/go-gin-demo/model/response"
	"github.com/akaedison/go-gin-demo/service"
	"github.com/gin-gonic/gin"
)

func CreateBook(c *gin.Context) {
	var book request.Book
	err := c.ShouldBindJSON(&book)
	if err != nil {
		response.CommonFailed("Bind Json Error", CodeBindJsonError, c)
		return
	}

	b := &model.Book{
		BookName:     book.BookName,
		Author:       book.Author,
		Price:        book.Price,
		Introduction: book.Introduction,
	}
	if err, bookAdded := service.CreateBook(b); err != nil {
		response.CommonFailed("Create Error", CodeDbErr, c)
		return
	} else {
		response.Created(bookAdded, "Create Book Success", c)
	}
}

func DeleteBook(c *gin.Context) {
	var reqId request.GetById
	if err := c.ShouldBindJSON(&reqId); err != nil {
		response.CommonFailed("Bind Error", CodeBindJsonError, c)
		return
	}

	if reqId.Id == 0 {
		response.CommonFailed("ID can not be null or 0", CodeCanNotBeNUll, c)
		return
	}

	if err := service.DeleteBook(reqId.Id); err != nil {
		response.CommonFailed("Delete Error", CodeDbErr, c)
	} else {
		response.DeleteSuccess(c)
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
