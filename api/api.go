package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetApiList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"teapot_url":          "https://api.akazwz.com/teapot",
		"token_url":           "https://api.akazwz.com/v1/token",
		"users_url":           "https://api.akazwz.com/v1/users",
		"users_authority_url": "https://api.akazwz.com/v1/users/authority",
		"users_password_url":  "https://api.akazwz.com/v1/users/password",
		"books_url":           "https://api.akazwz.com/v1/books",
		"file":                "https://api.akazwz.com/v1/file",
	})
}
