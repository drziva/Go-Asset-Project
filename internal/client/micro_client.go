package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type MicroClient struct {
	baseURL string
	http    *http.Client
}

func NewMicroClient(baseURL string) *MicroClient {
	return &MicroClient{
		baseURL: baseURL,
		http:    &http.Client{},
	}
}

func (c *MicroClient) GetHello(ctx context.Context) (string, error) {
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		fmt.Sprintf("%s/hello", c.baseURL),
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

	var msg string
	if err := json.NewDecoder(res.Body).Decode(&msg); err != nil {
		return "", err
	}

	return msg, err
}
