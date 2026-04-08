package service

import (
	"context"
	"go-project/internal/client"
)

type MicroService struct {
	microClient *client.MicroClient
}

func NewMicroService(client *client.MicroClient) *MicroService {
	return &MicroService{
		microClient: client,
	}
}

func (s *MicroService) GetHello(ctx context.Context) (string, error) {
	msg, err := s.microClient.GetHello(ctx)
	if err != nil {
		return "", err
	}

	return msg, nil
}
