package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		req := c.Request

		token := req.Header.Get("Authorization")

		if len(token) == 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status": http.StatusUnauthorized,
				"info":   "Unauthorized",
			})
			return
		}
		c.Next()
	}
}
