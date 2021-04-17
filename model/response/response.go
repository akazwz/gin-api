package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int         `json:"code,omitempty"`
	Data interface{} `json:"data,omitempty"` //omitempty nil or default
	Msg  string      `json:"msg,omitempty"`
}

const (
	SUCCESS = 0
	ERROR   = 7
)

func Unauthorized(code int, message string, c *gin.Context) {
	c.JSON(http.StatusUnauthorized, Response{
		Code: code,
		Msg:  message,
	})
}

func Forbidden(c *gin.Context) {
	c.JSON(http.StatusForbidden, Response{
		Msg: "Permission Denied",
	})
}

func DeleteSuccess(c *gin.Context) {
	c.JSON(http.StatusNoContent, nil)
}

func CommonFailed(message string, code int, c *gin.Context) {
	c.JSON(http.StatusBadRequest, Response{
		Code: code,
		Msg:  message,
	})
}

func CommonSuccess(code int, data interface{}, message string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Data: data,
		Msg:  message,
	})
}

func Created(data interface{}, message string, c *gin.Context) {
	c.JSON(http.StatusCreated, Response{
		Code: SUCCESS,
		Data: data,
		Msg:  message,
	})
}

func Result(code int, data interface{}, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

func Ok(c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, "Success", c)
}

func OkWithMessage(message string, c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, message, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(SUCCESS, data, "Success", c)
}

func OkWithDetail(data interface{}, message string, c *gin.Context) {
	Result(SUCCESS, data, message, c)
}

func Fail(c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, "Failed", c)
}

func FailWithMessage(message string, c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, message, c)
}

func FailWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(ERROR, data, message, c)
}
