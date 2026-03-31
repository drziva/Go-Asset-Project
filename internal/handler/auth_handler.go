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
	authService   *service.AuthService
	cookieService *service.CookieService
	accessTTL     int
}

func NewAuthHandler(authService *service.AuthService, cookieService *service.CookieService, accessTTL int) *AuthHandler {
	return &AuthHandler{
		authService,
		cookieService,
		accessTTL,
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

	if err := h.authService.SignUp(user); err != nil {
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

	user, token, err := h.authService.Login(dto)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid credentials",
		})

		return
	}

	h.cookieService.SetAccessTokenCookie(c, token, h.accessTTL)

	c.JSON(http.StatusAccepted, mappers.ToLoginResponse(*user))
}

func (h *AuthHandler) Me(c *gin.Context) {
	userID, exists := c.Get("user_id")

	if !exists {
		c.JSON(401, gin.H{
			"error": "unauthorized",
		})
		return
	}

	user, err := h.authService.Me(userID.(uint))

	if err != nil {
		c.JSON(404, gin.H{
			"error": "user not found",
		})
		return
	}

	c.JSON(http.StatusOK, mappers.ToLoginResponse(*user))
}
