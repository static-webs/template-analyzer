package clients

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LLMClient struct {
	httpClient http.Client
}

func NewLLMClient() *LLMClient {
	return &LLMClient{
		httpClient: http.Client{},
	}
}

func (c *LLMClient) Send(content string) (string, error) {
	payload, err := json.Marshal(map[string]any{
		"content": content,
	})
	if err != nil {
		return "", fmt.Errorf("failed to marshal request body: %w", err)
	}

	resp, err := c.httpClient.Post("http://localhost:8080/ai/respond", "application/json", bytes.NewReader(payload))
	if err != nil {
		return "", fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	return string(body), nil
}
