package commands

import (
	"github.com/danterolle/loqi/config"
	"github.com/danterolle/loqi/translate"
	"github.com/danterolle/loqi/translate/setup"
)

func SetupRun(cfg *config.Config, model string) (*translate.Core, func(), error) {
	return setup.SetupRun(cfg, model, logDiag, printBanner)
}
