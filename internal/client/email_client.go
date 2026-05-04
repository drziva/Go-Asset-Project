package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"go-project/internal/dto"
	"io"
	"net/http"
)

type EmailClient struct {
	baseURL string
	http    *http.Client
}

func NewEmailClient(baseURL string) *EmailClient {
	return &EmailClient{
		baseURL: baseURL,
		http:    &http.Client{},
	}
}

func (c *EmailClient) SendEmail(ctx context.Context) (string, error) { // Placeholder function
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		fmt.Sprintf("%s/email", c.baseURL),
		nil,
	)
	if err != nil {
		return "", err
	}

	res, err := c.http.Do(req)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()
	return "success", nil
}

func (c *EmailClient) SendVerificationEmail(ctx context.Context, emailRequest dto.SendEmailRequest) (string, error) {
	jsonBody, err := json.Marshal(emailRequest)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		fmt.Sprintf("%s/email/verification", c.baseURL),
		bytes.NewBuffer(jsonBody),
	)
	if err != nil {
		return "", err
	}

	res, err := c.http.Do(req)
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		body, _ := io.ReadAll(res.Body)
		return "", fmt.Errorf("email service error: %v", string(body))
	}
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	var resp dto.EmailResponse
	err = json.NewDecoder(res.Body).Decode(&resp)
	if err != nil {
		return "", err
	}

	return resp.Message, nil
}
