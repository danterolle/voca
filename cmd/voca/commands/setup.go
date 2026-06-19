package commands

import (
	"fmt"
	"time"

	"github.com/danterolle/voca/config"
	"github.com/danterolle/voca/translate"
	"github.com/danterolle/voca/translate/llamacpp"
	"github.com/danterolle/voca/translate/ollama"
)

func SetupRun(cfg *config.Config, model string) (*translate.Core, func(), error) {
	printBanner()

	switch cfg.Backend.Type {
	case "ollama":
		return setupOllama(cfg, model)
	case "llamacpp":
		return setupLlamaCpp(cfg, model)
	default:
		return nil, nil, fmt.Errorf("unsupported backend type: %q", cfg.Backend.Type)
	}
}

func setupOllama(cfg *config.Config, model string) (*translate.Core, func(), error) {
	ollamaCmd, started, err := SetupOllama(model, cfg.Backend.BaseURL)
	if err != nil {
		return nil, nil, err
	}

	var cleanup func()
	if started && ollamaCmd != nil {
		c := ollamaCmd
		cleanup = func() {
			ollama.UnloadModel(model, cfg.Backend.BaseURL)
			_ = c.Process.Kill()
		}
	} else {
		cleanup = func() { ollama.UnloadModel(model, cfg.Backend.BaseURL) }
	}

	backend := ollama.NewBackend(cfg.Backend.BaseURL, model, translate.NewDefaultPrompt())
	if np, ok := readFloatOption(cfg.Backend.Options, "num_predict"); ok {
		backend.NumPredict = int(np)
	}
	if to, ok := readFloatOption(cfg.Backend.Options, "timeout"); ok {
		backend.Client.Timeout = time.Duration(to) * time.Second
	}
	if t, ok := readFloatOption(cfg.Backend.Options, "temperature"); ok {
		backend.Temperature = t
	}
	if p, ok := readFloatOption(cfg.Backend.Options, "top_p"); ok {
		backend.TopP = p
	}

	return translate.NewCore(backend, translate.NewStaticLanguages()), cleanup, nil
}

func setupLlamaCpp(cfg *config.Config, model string) (*translate.Core, func(), error) {
	llamaCmd, started, err := SetupLlamaCpp(model, cfg.Backend.BaseURL, cfg.Backend.ModelPath, cfg.Backend.ServerArgs)
	if err != nil {
		return nil, nil, err
	}

	var cleanup func()
	if started && llamaCmd != nil {
		c := llamaCmd
		cleanup = func() { _ = c.Process.Kill() }
	} else {
		cleanup = func() {}
	}

	backend := llamacpp.NewBackend(cfg.Backend.BaseURL, model, translate.NewDefaultPrompt())
	if mt, ok := readFloatOption(cfg.Backend.Options, "num_predict"); ok {
		backend.MaxTokens = int(mt)
	}
	if to, ok := readFloatOption(cfg.Backend.Options, "timeout"); ok {
		backend.Client.Timeout = time.Duration(to) * time.Second
	}
	if t, ok := readFloatOption(cfg.Backend.Options, "temperature"); ok {
		backend.Temperature = t
	}
	if p, ok := readFloatOption(cfg.Backend.Options, "top_p"); ok {
		backend.TopP = p
	}

	return translate.NewCore(backend, translate.NewStaticLanguages()), cleanup, nil
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
