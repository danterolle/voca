package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/danterolle/voca/translate"
	"github.com/danterolle/voca/translate/ollama"
	"github.com/danterolle/voca/tui"
)

func newCore(model string) *translate.Core {
	prompt := translate.NewDefaultPrompt()
	return translate.NewCore(
		ollama.NewBackend("http://localhost:11434", model, prompt),
		prompt,
		translate.NewStaticLanguages(),
		model,
	)
}

func runTranslate(args []string) {
	fs := flag.NewFlagSet("translate", flag.ExitOnError)
	from := fs.String("from", "auto", "source language code")
	to := fs.String("to", "en", "target language code")
	model := fs.String("model", translate.DefaultModel, "Ollama model")
	h := fs.Bool("h", false, "show help")
	help := fs.Bool("help", false, "show help")
	fs.Parse(args)

	if *h || *help {
		printBanner()
		fmt.Println("Usage: voca translate [flags] <text|file>")
		fmt.Println()
		fs.PrintDefaults()
		fmt.Println()
		fmt.Println("Examples:")
		fmt.Println(`  voca translate --from it --to en "Ciao mondo!"`)
		fmt.Println("  voca translate --from en --to fr < README.md")
		os.Exit(0)
	}

	text := readInput(fs.Args())
	if text == "" {
		fmt.Fprintf(os.Stderr, "Usage: voca translate --from <lang> --to <lang> [text|file|stdin]\n")
		fs.PrintDefaults()
		os.Exit(1)
	}

	printBanner()
	ollamaCmd, started := setupOllama(*model)
	if started && ollamaCmd != nil {
		defer ollamaCmd.Process.Kill()
	}

	core := newCore(*model)
	ui := tui.NewCLIUI(*from, *to, text)
	if err := ui.Run(context.Background(), core); err != nil {
		fmt.Fprintf(os.Stderr, "  ✖ Error: %v\n", err)
		os.Exit(1)
	}
}

func runBatch(args []string) {
	fs := flag.NewFlagSet("batch", flag.ExitOnError)
	from := fs.String("from", "auto", "source language code")
	to := fs.String("to", "en", "target language code")
	model := fs.String("model", translate.DefaultModel, "Ollama model")
	h := fs.Bool("h", false, "show help")
	help := fs.Bool("help", false, "show help")
	fs.Parse(args)

	if *h || *help {
		printBanner()
		fmt.Println("Usage: voca batch [flags] [file]")
		fmt.Println()
		fs.PrintDefaults()
		fmt.Println()
		fmt.Println("Examples:")
		fmt.Println(`  voca batch --from en --to it < locales/en.json`)
		fmt.Println(`  voca batch --from en --to it locales/en.json`)
		fmt.Println(`  voca batch --from en --to fr README.md`)
		fmt.Println(`  echo "Hello world" | voca batch --from en --to it`)
		os.Exit(0)
	}

	input, err := readStdinOrFile(fs.Args())
	if err != nil || len(input) == 0 {
		fmt.Fprintf(os.Stderr, "Usage: voca batch --from <lang> --to <lang> [file]\n")
		fs.PrintDefaults()
		os.Exit(1)
	}

	ollamaCmd, started := setupOllama(*model)
	if started && ollamaCmd != nil {
		defer ollamaCmd.Process.Kill()
	}
	core := newCore(*model)
	ctx := context.Background()

	output, err := translate.Batch(ctx, core, input, *from, *to)
	if err != nil {
		fmt.Fprintf(os.Stderr, "  ✖ Error: %v\n", err)
		os.Exit(1)
	}

	os.Stdout.Write(output)
	fmt.Println()
}

func runTUI() {
	model := flag.String("model", translate.DefaultModel, "Ollama model to use for translation")
	flag.Parse()

	printBanner()
	ollamaCmd, started := setupOllama(*model)
	if started && ollamaCmd != nil {
		defer ollamaCmd.Process.Kill()
	}

	fmt.Printf("\n  Starting terminal interface...")
	time.Sleep(800 * time.Millisecond)
	fmt.Printf("\n")

	core := newCore(*model)
	ui := tui.NewBubbleTeaUI()
	if err := ui.Run(context.Background(), core); err != nil {
		fmt.Fprintf(os.Stderr, "  ✖ Error: %v\n", err)
	}
}
