package commands

import (
	"flag"
	"fmt"
	"os"

	"github.com/danterolle/voca/config"
)

var Version string

const defaultFrom = "auto"
const defaultTo = "en"

func Run(cfg *config.Config, args []string) {
	var err error
	if len(args) > 1 {
		switch args[1] {
		case "translate":
			err = RunTranslate(cfg, args[2:])
		case "batch":
			err = RunBatch(cfg, args[2:])
		case "-h", "--help":
			PrintUsage()
			return
		default:
			RunTUI(cfg, args[1:])
			return
		}
	} else {
		RunTUI(cfg, args[1:])
		return
	}
	if err != nil {
		Fatal(err)
	}
}

func PrintUsage() {
	printBanner()
	fmt.Println("Usage:")
	fmt.Println("  voca					   Start the terminal UI (default)")
	fmt.Println("  voca translate [flags] <text|file>           One-shot translation")
	fmt.Println("  voca batch [flags] <file|stdin>              Batch translate JSON or text")
	fmt.Println()
	fmt.Println("Global flags:")
	fmt.Println("  --config <path>                   Path to config file (optional)")
	fmt.Println("  -h, --help                        Show this help message")
	fmt.Println()
	cfg := config.Default()
	fmt.Println("Configurable flags (translate/batch):")
	fs := flag.NewFlagSet("translate", flag.ExitOnError)
	fs.String("from", defaultFrom, "source language code")
	fs.String("to", defaultTo, "target language code")
	fs.String("model", cfg.Backend.Model, "translation model")
	fs.PrintDefaults()
}

func Fatal(err error) {
	fmt.Fprintf(os.Stderr, "  ✖ Error: %v\n", err)
	os.Exit(1)
}

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
