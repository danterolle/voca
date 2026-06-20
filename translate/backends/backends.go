package backends

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/danterolle/voca/translate"
	"github.com/danterolle/voca/translate/llamacpp"
	"github.com/danterolle/voca/translate/ollama"
)

func NewBackend(backendType, baseURL, model string, options map[string]any, prompt translate.PromptBuilder) (translate.Backend, error) {
	switch backendType {
	case "ollama":
		b := ollama.NewBackend(baseURL, model, prompt)
		b.NumPredict = intOption(options, "num_predict", 2048)
		b.Client.Timeout = durationOption(options, "timeout", 2*time.Minute)
		b.Temperature = floatOption(options, "temperature", 0.0)
		b.TopP = floatOption(options, "top_p", 1.0)
		return b, nil
	case "llamacpp":
		b := llamacpp.NewBackend(baseURL, model, prompt)
		b.MaxTokens = intOption(options, "num_predict", 2048)
		b.Client.Timeout = durationOption(options, "timeout", 2*time.Minute)
		b.Temperature = floatOption(options, "temperature", 0.0)
		b.TopP = floatOption(options, "top_p", 1.0)
		return b, nil
	default:
		return nil, fmt.Errorf("unsupported backend type: %q", backendType)
	}
}

func UnloadBackend(backendType, model, baseURL string) {
	if backendType == "ollama" {
		var buf strings.Builder
		body := map[string]string{"model": model, "keep_alive": "0m", "unload": "true"}
		json.NewEncoder(&buf).Encode(body)
		client := &http.Client{Timeout: 30 * time.Second}
		client.Post(baseURL+"/api/generate", "application/json", strings.NewReader(buf.String()))
	}
}

func readFloatOption(options map[string]any, key string) (float64, bool) {
	v, ok := options[key]
	if !ok {
		return 0, false
	}
	switch n := v.(type) {
	case float64:
		return n, true
	case int:
		return float64(n), true
	}
	return 0, false
}

func intOption(options map[string]any, key string, defaultVal int) int {
	v, ok := readFloatOption(options, key)
	if !ok {
		return defaultVal
	}
	return int(v)
}

func floatOption(options map[string]any, key string, defaultVal float64) float64 {
	v, ok := readFloatOption(options, key)
	if !ok {
		return defaultVal
	}
	return v
}

func durationOption(options map[string]any, key string, defaultVal time.Duration) time.Duration {
	v, ok := readFloatOption(options, key)
	if !ok {
		return defaultVal
	}
	return time.Duration(v) * time.Second
}
