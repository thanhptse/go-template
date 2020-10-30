package handler

import "github.com/gin-gonic/gin"

func (h *Handler) ping() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	}
}
