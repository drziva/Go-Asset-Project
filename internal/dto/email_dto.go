package dto

type RequestType string

const (
	ResetPassword RequestType = "reset_password"
	ConfirmEmail  RequestType = "confirm_email"
	LinkAccount   RequestType = "link_account"
)

type SendEmailRequest struct {
	Email       string      `json:"email"`
	RequestType RequestType `json:"request_type"`
	Code        interface{} `json:"code"`
}

type EmailResponse struct {
	Message string `json:"message"`
}
