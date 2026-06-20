package setup

import (
	"fmt"
	"os/exec"

	"github.com/danterolle/loqi/config"
	"github.com/danterolle/loqi/translate"
)

type DiagFunc func(format string, args ...any)

func SetupRun(cfg *config.Config, model string, diag DiagFunc, banner func()) (*translate.Translator, func(), error) {
	if banner != nil {
		banner()
	}

	var (
		serverStarter func() (*exec.Cmd, bool, error)
		backendType   string
		unloadOnClose bool
	)

	switch cfg.Backend.Type {
	case "ollama":
		serverStarter = func() (*exec.Cmd, bool, error) {
			return SetupOllama(model, cfg.Backend.BaseURL, diag)
		}
		backendType = "ollama"
		unloadOnClose = true
	case "llamacpp":
		serverStarter = func() (*exec.Cmd, bool, error) {
			return SetupLlamaCpp(model, cfg.Backend.BaseURL, cfg.Backend.ModelPath, cfg.Backend.ServerArgs, diag)
		}
		backendType = "llamacpp"
		unloadOnClose = false
	default:
		return nil, nil, fmt.Errorf("unsupported backend type: %q", cfg.Backend.Type)
	}

	serverCmd, started, err := serverStarter()
	if err != nil {
		return nil, nil, err
	}

	var cleanup func()
	if started && serverCmd != nil {
		c := serverCmd
		cleanup = func() {
			if unloadOnClose {
				translate.UnloadBackend(backendType, model, cfg.Backend.BaseURL)
			}
			StopProcess(c)
		}
	} else if unloadOnClose {
		cleanup = func() { translate.UnloadBackend(backendType, model, cfg.Backend.BaseURL) }
	} else {
		cleanup = func() {}
	}

	backend, err := translate.NewBackend(backendType, cfg.Backend.BaseURL, model, cfg.Backend.Options, translate.NewChatPrompt())
	if err != nil {
		return nil, nil, err
	}

	return translate.NewTranslator(backend, translate.NewStaticLanguages()), cleanup, nil
}
