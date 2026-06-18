package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/danterolle/voca/config"
	"github.com/danterolle/voca/translate"
	"github.com/danterolle/voca/translate/ollama"
)

const defaultFrom = "auto"
const defaultTo = "en"

func parseTranslateFlags(name string, args []string, defaultModel string) (model, from, to string, fs *flag.FlagSet, h, help *bool) {
	model = defaultModel
	from = defaultFrom
	to = defaultTo

	fs = flag.NewFlagSet(name, flag.ExitOnError)
	fs.StringVar(&model, "model", model, "translation model")
	fs.StringVar(&from, "from", from, "source language code")
	fs.StringVar(&to, "to", to, "target language code")
	h = fs.Bool("h", false, "show help")
	help = fs.Bool("help", false, "show help")
	fs.Parse(args)
	return
}

func readFloat64Option(opts map[string]any, key string) (float64, bool) {
	v, ok := opts[key]
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

func newCore(cfg *config.Config, model string) (*translate.Core, error) {
	prompt := translate.NewDefaultPrompt()

	var backend *ollama.Backend
	switch cfg.Backend.Type {
	case "ollama":
		backend = ollama.NewBackend(cfg.Backend.BaseURL, model, prompt)
	default:
		return nil, fmt.Errorf("unsupported backend type: %q", cfg.Backend.Type)
	}

	if np, ok := readFloat64Option(cfg.Backend.Options, "num_predict"); ok {
		backend.NumPredict = int(np)
	}
	if to, ok := readFloat64Option(cfg.Backend.Options, "timeout"); ok {
		backend.Client.Timeout = time.Duration(to) * time.Second
	}
	if t, ok := readFloat64Option(cfg.Backend.Options, "temperature"); ok {
		backend.Temperature = t
	}
	if p, ok := readFloat64Option(cfg.Backend.Options, "top_p"); ok {
		backend.TopP = p
	}
	return translate.NewCore(backend, translate.NewStaticLanguages()), nil
}

func setupRun(cfg *config.Config, model string) (*translate.Core, func()) {
	printBanner()
	ollamaCmd, started, err := setupOllama(model)
	if err != nil {
		fatal(err)
	}

	var cleanup func()
	if started && ollamaCmd != nil {
		c := ollamaCmd
		cleanup = func() { _ = c.Process.Kill() }
	} else {
		cleanup = func() {}
	}

	core, err := newCore(cfg, model)
	if err != nil {
		cleanup()
		fatal(err)
	}

	return core, cleanup
}
