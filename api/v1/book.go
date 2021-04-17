package v1

import (
	_ "github.com/akaedison/go-gin-demo/docs"
	"github.com/akaedison/go-gin-demo/model"
	"github.com/akaedison/go-gin-demo/model/request"
	"github.com/akaedison/go-gin-demo/model/response"
	"github.com/akaedison/go-gin-demo/service"
	"github.com/gin-gonic/gin"
)

// CreateBook
// @Summary Create A Book
// @Title Create Book
// @Author zwz
// @Description create book
// @Tag book
// @Accept json
// @Produce json
// @Param book body request.Book true "json"
// @Param token header string true "token"
// @Success 201 {object} model.Book
// @Failure 400,401 {object} response.Response
// @Router /book [post]
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

// DeleteBook
// @Summary Delete A Book
// @Title Delete Book
// @Author zwz
// @Description delete book
// @Tag book
// @Accept json
// @Produce json
// @Param reqId body request.GetById true "id:2"
// @Param token header string true "token"
// @Success 204
// @Failure 400,401 {object} response.Response
// @Router /book [delete]
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
