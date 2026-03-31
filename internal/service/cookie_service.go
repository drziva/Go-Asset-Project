package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CookieService struct {
	cookieName   string
	cookieDomain string
	secure       bool
}

func NewCookieService(cookieName, cookieDomain string, secure bool) *CookieService {
	return &CookieService{
		cookieName,
		cookieDomain,
		secure,
	}
}

func (s *CookieService) SetAccessTokenCookie(c *gin.Context, token string, maxAgeSeconds int) {
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     s.cookieName,
		Value:    token,
		Path:     "/",
		Domain:   s.cookieDomain,
		MaxAge:   maxAgeSeconds,
		HttpOnly: true,
		Secure:   s.secure,
		SameSite: http.SameSiteLaxMode,
	})
}

func (s *CookieService) ClearAccessTokenCookie(c *gin.Context, token string, maxAgeSeconds int) {
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     s.cookieName,
		Value:    token,
		Path:     "/",
		Domain:   s.cookieDomain,
		MaxAge:   -1,
		HttpOnly: true,
		Secure:   s.secure,
		SameSite: http.SameSiteLaxMode,
	})
}
