package ollama

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	httpclient "github.com/danterolle/loqi/translate/http"
)

type chatRequest struct {
	Model    string               `json:"model"`
	Messages []httpclient.Message `json:"messages"`
	Stream   bool                 `json:"stream"`
	Options  map[string]any       `json:"options"`
}

type Backend struct {
	Config httpclient.BackendConfig
}

func NewBackend(config httpclient.BackendConfig) *Backend {
	if config.Client == nil {
		config.Client = httpclient.NewHTTPClient()
	}
	return &Backend{Config: config}
}

func (b *Backend) Translate(ctx context.Context, text, source, target string) (string, error) {
	if strings.TrimSpace(text) == "" {
		return "", nil
	}
	if source == target {
		return text, nil
	}

	return httpclient.PostJSON(ctx, b.Config.Client, b.Config.BaseURL+"/api/chat", "ollama",
		b.buildRequestBody(text, source, target),
		func(data []byte) (string, error) {
			var cr struct {
				Message httpclient.Message `json:"message"`
			}
			if err := json.Unmarshal(data, &cr); err != nil {
				return "", fmt.Errorf("ollama: decode: %w", err)
			}
			return strings.TrimSpace(cr.Message.Content), nil
		})
}

func (b *Backend) buildRequestBody(text, source, target string) chatRequest {
	options := map[string]any{
		"temperature": b.Config.Temperature,
		"top_p":       b.Config.TopP,
	}
	if b.Config.MaxTokens > 0 {
		options["num_predict"] = b.Config.MaxTokens
	}

	return chatRequest{
		Model: b.Config.Model,
		Messages: []httpclient.Message{
			{Role: "system", Content: b.Config.Prompt.System()},
			{Role: "user", Content: b.Config.Prompt.Translate(text, source, target)},
		},
		Stream:  false,
		Options: options,
	}
}
