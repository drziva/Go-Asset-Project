package service

import (
	"go-project/internal/constants"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CookieService struct {
	cookieDomain string
	secure       bool
}

func NewCookieService(cookieDomain string, secure bool) *CookieService {
	return &CookieService{
		cookieDomain,
		secure,
	}
}

func (s *CookieService) SetAccessTokenCookie(c *gin.Context, token string, maxAgeSeconds int) {
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     constants.AccessCookieName,
		Value:    token,
		Path:     "/",
		Domain:   s.cookieDomain,
		MaxAge:   maxAgeSeconds,
		HttpOnly: true,
		Secure:   s.secure,
		SameSite: http.SameSiteLaxMode,
	})
}

func (s *CookieService) ClearAccessTokenCookie(c *gin.Context) {
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     constants.AccessCookieName,
		Value:    "",
		Path:     "/",
		Domain:   s.cookieDomain,
		MaxAge:   -1,
		HttpOnly: true,
		Secure:   s.secure,
		SameSite: http.SameSiteLaxMode,
	})
}
