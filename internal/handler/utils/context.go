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

func ExtractIDParam(c *gin.Context) (uint, error) {
	idParam := c.Param("id") // string

	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		return 0, err
	}

	return uint(id), err
}
