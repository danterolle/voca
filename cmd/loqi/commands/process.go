package commands

import (
	"os/exec"

	"github.com/danterolle/loqi/translate/setup"
)

func stopProcess(cmd *exec.Cmd) {
	setup.StopProcess(cmd)
}
