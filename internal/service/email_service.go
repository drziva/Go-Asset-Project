package service

import (
	"context"
	"fmt"
	"go-project/internal/client"
	"go-project/internal/dto"
	"go-project/internal/errors"
)

type EmailService struct {
	emailClient *client.EmailClient
}

func NewEmailService(client *client.EmailClient) *EmailService {
	return &EmailService{
		emailClient: client,
	}
}

func (s *EmailService) SendEmail(ctx context.Context) (string, error) {
	msg, err := s.emailClient.SendEmail(ctx)
	if err != nil {
		return "", err
	}

	return msg, nil
}

func (s *EmailService) SendVerificationEmail(ctx context.Context, emailRequest dto.SendEmailRequest) (string, error) {
	msg, err := s.emailClient.SendVerificationEmail(ctx, emailRequest)
	if err != nil {
		fmt.Printf("email service error: %v", err)
		return "", fmt.Errorf("%w: %v", errors.ErrEmailServiceFailed, err)
	}

	return msg, nil
}
