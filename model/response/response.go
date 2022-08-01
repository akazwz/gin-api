package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code,omitempty"`
	Data interface{} `json:"data,omitempty"`
	Msg  string      `json:"msg,omitempty"`
}

const (
	SUCCESS = 2000
	ERROR   = 4000
)

func Ok(code int, data interface{}, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}

func Created(code int, data interface{}, msg string, c *gin.Context) {
	c.JSON(http.StatusCreated, Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}

func BadRequest(code int, data interface{}, msg string, c *gin.Context) {
	c.JSON(http.StatusBadRequest, Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}

func Unauthorized(code int, msg string, c *gin.Context) {
	c.JSON(http.StatusUnauthorized, Response{
		Code: code,
		Msg:  msg,
	})
}

func Forbidden(code int, msg string, c *gin.Context) {
	c.JSON(http.StatusForbidden, Response{
		Code: code,
		Msg:  msg,
	})
}

func NotFound(code int, msg string, c *gin.Context) {
	c.JSON(http.StatusNotFound, Response{
		Code: code,
		Msg:  msg,
	})
}
