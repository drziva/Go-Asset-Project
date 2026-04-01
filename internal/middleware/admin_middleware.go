package middleware

import (
	"go-project/internal/constants"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		isAdmin, exists := c.Get(constants.IsAdminKey)
		if !exists || isAdmin.(bool) == false {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "forbidden",
			})

			return
		}

		c.Next()
	}
}
