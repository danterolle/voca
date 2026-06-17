package translate

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type OllamaBackend struct {
	BaseURL string
	Model   string
	Prompt  PromptBuilder
}

func NewOllamaBackend(baseURL, model string, prompt PromptBuilder) *OllamaBackend {
	return &OllamaBackend{
		BaseURL: baseURL,
		Model:   model,
		Prompt:  prompt,
	}
}

func (b *OllamaBackend) Translate(ctx context.Context, text, source, target string) (string, error) {
	if strings.TrimSpace(text) == "" {
		return "", nil
	}
	if source == target {
		return text, nil
	}

	body := chatRequest{
		Model: b.Model,
		Messages: []message{
			{Role: "system", Content: b.Prompt.System()},
			{Role: "user", Content: b.Prompt.Translate(text, source, target)},
		},
		Stream: false,
		Options: map[string]any{
			"temperature": 0.0,
			"num_predict": 512,
			"top_p":       1.0,
		},
	}

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(body); err != nil {
		return "", fmt.Errorf("encode: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, b.BaseURL+"/api/chat", &buf)
	if err != nil {
		return "", fmt.Errorf("request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("ollama: %w", err)
	}
	defer resp.Body.Close()

	var cr chatResponse
	if err := json.NewDecoder(resp.Body).Decode(&cr); err != nil {
		return "", fmt.Errorf("decode: %w", err)
	}

	return strings.TrimSpace(cr.Message.Content), nil
}
