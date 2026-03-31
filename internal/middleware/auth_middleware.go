package middleware

import (
	"go-project/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(jwtService *service.JWTService, cookieName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie(cookieName)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "missing authorization cookie",
			})

			return
		}

		claims, err := jwtService.ValidateAccessToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "invalid or expired token",
			})

			return
		}

		c.Set("user_id", claims.ID)
		c.Set("is_admin", claims.IsAdmin)

		c.Next()
	}
}
