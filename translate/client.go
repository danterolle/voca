package translate

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type chatRequest struct {
	Model    string         `json:"model"`
	Messages []message      `json:"messages"`
	Stream   bool           `json:"stream"`
	Options  map[string]any `json:"options"`
}

type chatResponse struct {
	Message message `json:"message"`
}

func buildPrompt(text, source, target string) string {
	src := Languages[source]
	if source == "auto" {
		src = "Auto detect"
	}
	return fmt.Sprintf("Translate from %s to %s:\n\n%s", src, Languages[target], text)
}

func Translate(text, source, target, model string) (string, error) {
	if strings.TrimSpace(text) == "" {
		return "", nil
	}
	if source == target {
		return text, nil
	}

	body := chatRequest{
		Model: model,
		Messages: []message{
			{Role: "system", Content: "You are a translator. Translate the user's text accurately. Preserve meaning, tone, and sentence structure. Output only the translation — no greetings, explanations, or commentary."},
			{Role: "user", Content: buildPrompt(text, source, target)},
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

	resp, err := http.Post("http://localhost:11434/api/chat", "application/json", &buf)
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
