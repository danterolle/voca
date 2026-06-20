package commands

import (
	"os/exec"

	"github.com/danterolle/loqi/translate/setup"
)

func SetupLlamaCpp(model, baseURL, modelPath string, serverArgs []string) (cmd *exec.Cmd, started bool, err error) {
	return setup.SetupLlamaCpp(model, baseURL, modelPath, serverArgs, logDiag)
}
