package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/danterolle/voca/translate"
)

var Version string

func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "translate":
			runTranslate(os.Args[2:])
			return
		case "batch":
			runBatch(os.Args[2:])
			return
		case "-h", "--help":
			printUsage()
			return
		}
	}
	runTUI()
}

func printUsage() {
	printBanner()
	fmt.Println("Usage:")
	fmt.Println("  voca                   Start the terminal UI (default)")
	fmt.Println("  voca translate [flags] <text|file>   One-shot translation")
	fmt.Println("  voca batch [flags] <file|stdin>      Batch translate JSON or text")
	fmt.Println()
	fmt.Println("Flags:")
	fmt.Println("  -h, --help            Show this help message")
	fmt.Println()
	fmt.Println("Translate subcommand flags:")
	fs := flag.NewFlagSet("translate", flag.ExitOnError)
	fs.String("from", "auto", "source language code")
	fs.String("to", "en", "target language code")
	fs.String("model", translate.DefaultModel, "Ollama model")
	fs.PrintDefaults()
	fmt.Println()
	fmt.Println("Batch subcommand flags:")
	bs := flag.NewFlagSet("batch", flag.ExitOnError)
	bs.String("from", "auto", "source language code")
	bs.String("to", "en", "target language code")
	bs.String("model", translate.DefaultModel, "Ollama model")
	bs.PrintDefaults()
}
