package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PermissionMiddleware(permission string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// sementara biar test lolos
		auth := c.GetHeader("Authorization")
		if auth == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// lanjutkan request
		c.Next()
	}
}