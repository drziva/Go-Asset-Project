package handler

import (
	"go-project/internal/dto"
	handler "go-project/internal/handler/errors"
	"go-project/internal/mappers"
	"go-project/internal/models"
	"go-project/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service *service.AuthService
}

func NewAuthHandler(service *service.AuthService) *AuthHandler {
	return &AuthHandler{
		service,
	}
}

func (h *AuthHandler) SignUp(c *gin.Context) {
	var dto dto.SignUpDTO

	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	user := &models.User{
		Email:    dto.Email,
		Name:     dto.Name,
		Password: dto.Password,
	}

	if err := h.service.SignUp(user); err != nil {
		handler.HandleError(c, err)

		return
	}

	c.JSON(http.StatusCreated, mappers.ToLoginResponse(*user))
}

func (h *AuthHandler) Login(c *gin.Context) {
	var dto dto.LoginDTO

	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	user, err := h.service.Login(dto)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid credentials",
		})

		return
	}

	c.JSON(http.StatusAccepted, mappers.ToLoginResponse(*user))
}
