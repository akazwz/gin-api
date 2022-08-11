package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Healthz(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func Endpoints(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"teapot":   "/teapot",
		"file":     "/file",
		"auth":     "/auth",
		"s3":       "/s3",
		"image":    "/image",
		"posts":    "/posts",
		"projects": "/project",
		"healthz":  "/healthz",
	})
}

func Teapot(c *gin.Context) {
	c.JSON(http.StatusTeapot, gin.H{
		"message": "I'm a teapot",
		"story": "This code was defined in 1998 " +
			"as one of the traditional IETF April Fools' jokes," +
			" in RFC 2324, Hyper Text Coffee Pot Control Protocol," +
			" and is not expected to be implemented by actual HTTP servers." +
			" However, known implementations do exist.",
	})
}

func NotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"message": "Not Found",
	})
}
