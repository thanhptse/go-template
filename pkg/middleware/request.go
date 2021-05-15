package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// SetRequestID ...
func SetRequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		reqHeaderName := "X-Request-ID"
		requestID := c.Request.Header.Get(reqHeaderName)
		if requestID == "" {
			requestID = uuid.NewString()
		}
		c.Set("RequestID", requestID)
		c.Writer.Header().Set(reqHeaderName, requestID)
	}
}
