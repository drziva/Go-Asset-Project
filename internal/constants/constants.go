package constants

type AuthProvider string

const (
	UserIDKey                       = "user_id"
	IsAdminKey                      = "is_admin"
	AccessCookieName                = "accessToken"
	AuthProviderGoogle AuthProvider = "google"
	AuthProviderGithub AuthProvider = "github"
	AuthProviderLocal  AuthProvider = "local"
)
