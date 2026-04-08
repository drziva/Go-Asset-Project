package utils

import (
	"go-project/internal/constants"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ExtractUserID(c *gin.Context) uint {
	userID, _ := c.Get(constants.UserIDKey)

	return userID.(uint)
}

func ExtractIsAdmin(c *gin.Context) bool {
	isAdmin, _ := c.Get(constants.IsAdminKey)

	return isAdmin.(bool)
}

func ExtractIDParam(c *gin.Context) (uint, error) {
	idParam := c.Param("id") // string

	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		return 0, err
	}

	return uint(id), err
}
