package handler

import (
	errorHandler "go-project/internal/handler/errors"
	"go-project/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MicroHandler struct {
	service *service.MicroService
}

func NewMicroHandler(service *service.MicroService) *MicroHandler {
	return &MicroHandler{
		service,
	}
}

func (h *MicroHandler) GetHello(c *gin.Context) {
	msg, err := h.service.GetHello(c.Request.Context())

	if err != nil {
		errorHandler.HandleError(c, err)

		return
	}

	c.JSON(http.StatusOK, msg)
}
