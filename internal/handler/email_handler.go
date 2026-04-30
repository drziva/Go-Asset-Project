package handler

import (
	"go-project/internal/dto"
	apiErrors "go-project/internal/handler/errors"
	errorHandler "go-project/internal/handler/errors"
	"go-project/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type EmailHandler struct {
	service *service.EmailService
}

func NewEmailHandler(service *service.EmailService) *EmailHandler {
	return &EmailHandler{
		service,
	}
}

func (h *EmailHandler) SendEmail(c *gin.Context) {
	msg, err := h.service.SendEmail(c.Request.Context())

	if err != nil {
		errorHandler.HandleError(c, err)

		return
	}

	c.JSON(http.StatusOK, msg)
}

func (h *EmailHandler) SendVerificationEmail(c *gin.Context) {
	ctx := c.Request.Context()
	var emailRequest dto.SendEmailRequest
	if err := c.ShouldBindJSON(&emailRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "missing or invalid email request object",
		})
		return
	}

	msg, err := h.service.SendVerificationEmail(ctx, emailRequest)
	if err != nil {
		apiErrors.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, msg)
}
