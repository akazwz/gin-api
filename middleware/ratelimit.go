package middleware

import (
	"time"

	"github.com/didip/tollbooth/v7"
	"github.com/didip/tollbooth/v7/limiter"
	"github.com/gin-gonic/gin"
)

func LimitByRequest(count float64) gin.HandlerFunc {
	opt := limiter.ExpirableOptions{
		DefaultExpirationTTL: time.Hour,
	}
	lmt := tollbooth.NewLimiter(count, &opt)
	return func(c *gin.Context) {
		httpErr := tollbooth.LimitByRequest(lmt, c.Writer, c.Request)
		if httpErr != nil {
			c.Data(httpErr.StatusCode, lmt.GetMessageContentType(), []byte(httpErr.Message))
			c.Abort()
		} else {
			c.Next()
		}
	}
}
