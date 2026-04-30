package errors

import "errors"

var (
	// generic resource errors
	ErrNotFound           = errors.New("resource not found")
	ErrEmailAlreadyExists = errors.New("user with this email already exists")
	ErrConflict           = errors.New("conflict")

	// validation / input errors
	ErrInvalidInput         = errors.New("invalid input")
	ErrMissingRequiredField = errors.New("missing required field")
	ErrInvalidFormat        = errors.New("invalid format")
	ErrInvalidCredentials   = errors.New("invalid credentials")
	ErrUnauthorized         = errors.New("unauthorized")
	ErrForbidden            = errors.New("forbidden")
	ErrInvalidLinkToken     = errors.New("invalid or expired link token")
	ErrInvalidReference     = errors.New("invalid reference")
	ErrUnverifiedEmail      = errors.New("email address not verified")
)
