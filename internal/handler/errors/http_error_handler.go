package handler

import (
	"errors"
	"log"
	"net/http"

	"go-project/internal/constants"
	appErrors "go-project/internal/errors"

	"github.com/gin-gonic/gin"
)

func HandleError(c *gin.Context, err error) {
	log.Printf("ERROR TYPE: %T\n", err)
	log.Printf("ERROR VALUE: %+v\n", err)

	switch {

	// 404
	case errors.Is(err, appErrors.ErrNotFound):

		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})

	// 409
	case errors.Is(err, appErrors.ErrEmailAlreadyExists),
		errors.Is(err, appErrors.ErrConflict):

		c.JSON(http.StatusConflict, gin.H{
			"error": err.Error(),
		})

	// 400
	case errors.Is(err, appErrors.ErrInvalidInput),
		errors.Is(err, appErrors.ErrMissingRequiredField),
		errors.Is(err, appErrors.ErrInvalidLinkToken),
		errors.Is(err, appErrors.ErrInvalidFormat):

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

	// 401
	case errors.Is(err, appErrors.ErrInvalidCredentials),
		errors.Is(err, appErrors.ErrInvalidVerificationCode),
		errors.Is(err, appErrors.ErrUnauthorized):

		if errors.Is(err, appErrors.ErrUnauthorized) {
			c.SetCookie(constants.AccessCookieName, "", -1, "/", "", true, true)
		}

		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})

	// 403
	case errors.Is(err, appErrors.ErrForbidden):

		c.JSON(http.StatusForbidden, gin.H{
			"error": err.Error(),
		})

	case errors.Is(err, appErrors.ErrEmailServiceFailed):
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "email service failed",
		})

	//500
	default:

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})

	}
}
