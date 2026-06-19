package commands

import (
	"context"
	"flag"
	"os/signal"
	"syscall"
	"time"

	"github.com/danterolle/voca/config"
	"github.com/danterolle/voca/tui"
)

func RunTUI(cfg *config.Config, args []string) error {
	model := cfg.Backend.Model
	fs := flag.NewFlagSet("tui", flag.ExitOnError)
	fs.StringVar(&model, "model", model, "translation model")
	fs.Parse(args)

	core, cleanup, err := SetupRun(cfg, model)
	if err != nil {
		return err
	}
	defer cleanup()

	logDiag("\n  Starting terminal interface...")
	if !Quiet {
		time.Sleep(800 * time.Millisecond)
	}
	logDiag("\n")

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	return tui.RunBubbleTea(ctx, core.Backend, core.Languages)
}
