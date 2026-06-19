package commands

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/danterolle/voca/translate/ollama"
)

func SetupOllama(model, baseURL string) (cmd *exec.Cmd, started bool, err error) {
	if _, err := exec.LookPath("ollama"); err != nil {
		return nil, false, fmt.Errorf("ollama not found — install from https://ollama.com")
	}

	if !ollama.Reachable(baseURL) {
		fmt.Fprintf(os.Stderr, "  ◆ Starting Ollama... ")
		cmd = exec.Command("ollama", "serve")
		if err := cmd.Start(); err != nil {
			return nil, false, fmt.Errorf("failed to start Ollama: %w", err)
		}
		started = true
		if !ollama.WaitForReady(30, baseURL) {
			if cmd.Process != nil {
				_ = cmd.Process.Kill()
			}
			return nil, started, fmt.Errorf("timeout waiting for Ollama to start")
		}
		fmt.Fprintf(os.Stderr, "online\n")
	}

	if !ollama.ModelExists(model, baseURL) {
		fmt.Fprintf(os.Stderr, "  ◆ Pulling %s...\n", model)
		if err := ollama.PullModel(model, baseURL); err != nil {
			if started && cmd != nil {
				_ = cmd.Process.Kill()
			}
			return nil, started, fmt.Errorf("pull failed: %w", err)
		}
		fmt.Fprintf(os.Stderr, "  ◆ Model ready\n")
	}

	return cmd, started, nil
}
