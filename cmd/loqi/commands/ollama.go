package commands

import (
	"os/exec"

	"github.com/danterolle/loqi/translate/setup"
)

func SetupOllama(model, baseURL string) (cmd *exec.Cmd, started bool, err error) {
	return setup.SetupOllama(model, baseURL, logDiag)
}
