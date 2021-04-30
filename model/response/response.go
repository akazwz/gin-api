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
	SUCCESS  = 2000
	ERROR    = 4000
	PROGRESS = 2020
)

func Unauthorized(code int, message string, c *gin.Context) {
	c.JSON(http.StatusUnauthorized, Response{
		Code: code,
		Msg:  message,
	})
}

func PermissionDenied(code int, message string, c *gin.Context) {
	c.JSON(http.StatusForbidden, Response{
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

func SuccessWithMessage(message string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: SUCCESS,
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

func Accepted(fileUploadStatus interface{}, c *gin.Context) {
	c.JSON(http.StatusAccepted, Response{
		Code: PROGRESS,
		Data: fileUploadStatus,
		Msg:  "File is Uploading",
	})
}

func Result(code int, data interface{}, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}
