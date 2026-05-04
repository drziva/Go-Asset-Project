package constants

import "go-project/internal/dto"

type AuthProvider string

const (
	UserIDKey                            = "user_id"
	IsAdminKey                           = "is_admin"
	AccessCookieName                     = "accessToken"
	AuthProviderGoogle   AuthProvider    = "google"
	AuthProviderGithub   AuthProvider    = "github"
	AuthProviderLocal    AuthProvider    = "local"
	LinkRequest          dto.RequestType = "link_request"
	ResetPasswordRequest dto.RequestType = "reset_password"
)
