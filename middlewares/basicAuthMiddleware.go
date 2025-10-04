package middlewares

import (
	"github.com/gin-gonic/gin"
)

func BasicAuthMiddleware() gin.HandlerFunc {
	return gin.BasicAuth(gin.Accounts{
		"admin": "admin123",
	})
}
