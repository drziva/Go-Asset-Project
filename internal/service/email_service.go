package service

import (
	"context"
	"go-project/internal/client"
	"go-project/internal/dto"
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
		return "", err
	}

	return msg, nil
}
