package handler

import (
	"go-project/internal/dto"
	apiErrors "go-project/internal/handler/errors"
	"go-project/internal/handler/utils"
	"go-project/internal/mappers"
	"go-project/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService   *service.AuthService
	emailService  *service.EmailService
	cookieService *service.CookieService
	accessTTL     int
}

func NewAuthHandler(authService *service.AuthService, emailService *service.EmailService, cookieService *service.CookieService, accessTTL int) *AuthHandler {
	return &AuthHandler{
		authService,
		emailService,
		cookieService,
		accessTTL,
	}
}

func (h *AuthHandler) SignUp(c *gin.Context) {
	var signUpDTO dto.SignUpDTO
	ctx := c.Request.Context()

	if err := c.ShouldBindJSON(&signUpDTO); err != nil {
		apiErrors.HandleError(c, err)

		return
	}

	authResult, err := h.authService.SignUp(signUpDTO)

	if err != nil {
		apiErrors.HandleError(c, err)
		return
	}

	if authResult != nil && authResult.RequiresLink == true && authResult.VerificationCode != "" {
		emailRequestDTO := &dto.SendEmailRequest{
			Email:       signUpDTO.Email,
			RequestType: "link_account",
			Code:        authResult.VerificationCode,
		}

		msg, err := h.emailService.SendVerificationEmail(ctx, *emailRequestDTO)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, msg)
		return
	}

	c.JSON(http.StatusCreated, mappers.ToLoginResponse(*authResult.User))
}

func (h *AuthHandler) Login(c *gin.Context) {
	var dto dto.LoginDTO

	if err := c.ShouldBindJSON(&dto); err != nil {
		apiErrors.HandleError(c, err)

		return
	}

	user, token, err := h.authService.Login(dto)
	if err != nil {
		apiErrors.HandleError(c, err)

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
		return
	}

	if authResult.RequiresLink == true {
		c.JSON(http.StatusOK, gin.H{"link_token": authResult.LinkToken})
		return
	}

	if err != nil {
		apiErrors.HandleError(c, err)
		return
	}

	h.cookieService.SetAccessTokenCookie(c, token, h.accessTTL)

	c.JSON(http.StatusOK, mappers.ToLoginResponse(*authResult.User))
}

func (h *AuthHandler) LinkAndLogin(c *gin.Context) {
	var dto dto.LinkRequest

	if err := c.ShouldBindJSON(&dto); err != nil {
		apiErrors.HandleError(c, err)
		return
	}
	user, token, err := h.authService.LinkAndLogin(dto)
	if err != nil {
		apiErrors.HandleError(c, err)
		return
	}

	h.cookieService.SetAccessTokenCookie(c, token, h.accessTTL)

	c.JSON(http.StatusOK, mappers.ToLoginResponse(*user))
}
