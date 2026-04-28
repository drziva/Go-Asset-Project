package handler

import (
	"fmt"
	"go-project/internal/dto"
	apiErrors "go-project/internal/handler/errors"
	httpErrors "go-project/internal/handler/errors"
	"go-project/internal/handler/utils"
	"go-project/internal/mappers"
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

	user, err := h.authService.SignUp(dto)
	if err != nil {
		httpErrors.HandleError(c, err)

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

	c.JSON(http.StatusOK, mappers.ToLoginResponse(*user))
}

func (h *AuthHandler) Logout(c *gin.Context) {
	h.cookieService.ClearAccessTokenCookie(c)

	c.JSON(http.StatusOK, gin.H{
		"message": "user has been logged out successfully",
	})
}

func (h *AuthHandler) Me(c *gin.Context) {
	userID := utils.ExtractUserID(c)
	user, err := h.authService.Me(userID)

	if err != nil {
		apiErrors.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, mappers.ToLoginResponse(*user))
}

func (h *AuthHandler) GoogleLogin(c *gin.Context) {
	url := h.authService.GetGoogleLoginURL()
	c.Redirect(302, url)
}

func (h *AuthHandler) GoogleCallback(c *gin.Context) {
	code := c.Query("code")

	authResult, token, err := h.authService.HandleGoogleCallback(c.Request.Context(), code)
	if err != nil {
		apiErrors.HandleError(c, err)
	}

	if authResult.Linked == false {
		c.JSON(http.StatusOK, gin.H{"link_token": authResult.LinkToken})
	}

	if err != nil {
		c.JSON(500, gin.H{"error": "auth failed"})
		return
	}

	h.cookieService.SetAccessTokenCookie(c, token, h.accessTTL)

	c.JSON(http.StatusOK, mappers.ToLoginResponse(*authResult.User))
}

func (h *AuthHandler) LinkAndLogin(c *gin.Context) {
	var dto dto.LinkRequest

	if err := c.ShouldBindJSON(&dto); err != nil {
		apiErrors.HandleError(c, err)
	}
	fmt.Print("-----------TOKEN------------------", dto.LinkToken)

	user, token, err := h.authService.LinkAndLogin(dto)
	if err != nil {
		apiErrors.HandleError(c, err)
		return
	}

	h.cookieService.SetAccessTokenCookie(c, token, h.accessTTL)

	c.JSON(http.StatusOK, mappers.ToLoginResponse(*user))
}
