package httputil

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client interface {
	SendRequest(method, url string, headers map[string]string, body any) ([]byte, int, error)
}

type client struct {
	httpClient *http.Client
}

func NewClient() Client {
	return &client{
		httpClient: &http.Client{},
	}
}

// SendRequest sends an HTTP request with JSON payload and returns the response body + status code
func (c *client) SendRequest(method, url string, headers map[string]string, body any) ([]byte, int, error) {
	var reqBody []byte
	var err error

	if body != nil {
		reqBody, err = json.Marshal(body)
		if err != nil {
			return nil, 0, fmt.Errorf("failed to marshal body: %v", err)
		}
	}

	fmt.Println("request body:", string(reqBody))

	req, err := http.NewRequest(method, url, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, 0, fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, 0, fmt.Errorf("request failed: %v", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, resp.StatusCode, fmt.Errorf("failed to read response: %v", err)
	}

	return respBody, resp.StatusCode, nil
}
