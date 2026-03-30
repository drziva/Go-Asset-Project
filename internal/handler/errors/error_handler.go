package handler

import (
	"errors"
	appErrors "go-project/internal/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleError(c *gin.Context, err error) {

	switch {

	case errors.Is(err, appErrors.ErrEmailAlreadyExists):
		c.JSON(http.StatusConflict, gin.H{
			"error": err.Error(),
		})

	case errors.Is(err, appErrors.ErrNotFound):
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})

	case errors.Is(err, appErrors.ErrInvalidCredentials):
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})

	case errors.Is(err, appErrors.ErrUnauthorized):
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})

	default:
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})

	}
}
