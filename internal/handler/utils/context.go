package utils

import (
	"go-project/internal/constants"

	"github.com/gin-gonic/gin"
)

func ExtractUserID(c *gin.Context) uint {
	userID, _ := c.Get(constants.UserIDKey)

	return userID.(uint)
}
