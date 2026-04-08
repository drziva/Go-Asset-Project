package middleware

import (
	"go-project/internal/constants"
	"go-project/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(jwtService *service.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie(constants.AccessCookieName)
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

		c.Set(constants.UserIDKey, uint(claims.ID))
		c.Set(constants.IsAdminKey, claims.IsAdmin)

		c.Next()
	}
}
