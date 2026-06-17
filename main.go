package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/danterolle/voca/translate"
	"github.com/danterolle/voca/translate/ollama"
	"github.com/danterolle/voca/tui"
)

var Version string

func main() {
	model := flag.String("model", translate.DefaultModel, "Ollama model to use for translation")
	flag.Parse()

	startedOllama := false
	var ollamaCmd *exec.Cmd

	printBanner()

	if _, err := exec.LookPath("ollama"); err != nil {
		fmt.Printf("  ✖ ollama not found. Install it from https://ollama.com\n")
		os.Exit(1)
	}

	if !ollama.Reachable() {
		fmt.Printf("  ◆ Starting Ollama... ")
		ollamaCmd = exec.Command("ollama", "serve")
		if err := ollamaCmd.Start(); err != nil {
			fmt.Printf("\n  ✖ Failed to start Ollama: %v\n", err)
			os.Exit(1)
		}
		startedOllama = true
		if !ollama.WaitForReady(30) {
			fmt.Printf("timeout waiting for Ollama to start\n")
			ollamaCmd.Process.Kill()
			os.Exit(1)
		}
		fmt.Printf("online\n")
	}

	if !ollama.ModelExists(*model) {
		fmt.Printf("  ◆ Pulling %s...\n", *model)
		if err := ollama.PullModel(*model); err != nil {
			fmt.Printf("  ✖ Pull failed: %v\n", err)
			if startedOllama && ollamaCmd != nil {
				ollamaCmd.Process.Kill()
			}
			os.Exit(1)
		}
		fmt.Printf("  ◆ Model ready\n")
	}

	fmt.Printf("\n  Starting terminal interface...")
	time.Sleep(800 * time.Millisecond)
	fmt.Printf("\n")

	core := translate.NewCore(
		ollama.NewBackend("http://localhost:11434", *model, translate.NewDefaultPrompt()),
		translate.NewDefaultPrompt(),
		translate.NewStaticLanguages(),
		*model,
	)

	ui := tui.NewBubbleTeaUI()
	if err := ui.Run(context.Background(), core); err != nil {
		fmt.Fprintf(os.Stderr, "  ✖ Error: %v\n", err)
	}

	if startedOllama && ollamaCmd != nil {
		ollamaCmd.Process.Kill()
	}
}

func printBanner() {
	gradient := []string{
		"\033[38;5;255m",
		"\033[38;5;230m",
		"\033[38;5;229m",
		"\033[38;5;221m",
		"\033[38;5;215m",
		"\033[38;5;203m",
	}
	reset := "\033[0m"

	lines := []string{
		"  ██╗   ██╗ ██████╗  ██████╗ █████╗ ",
		"  ██║   ██║██╔═══██╗██╔════╝██╔══██╗",
		"  ██║   ██║██║   ██║██║     ███████║",
		"  ╚██╗ ██╔╝██║   ██║██║     ██╔══██║",
		"   ╚████╔╝ ╚██████╔╝╚██████╗██║  ██║",
		"    ╚═══╝   ╚═════╝  ╚═════╝╚═╝  ╚═╝",
	}

	fmt.Println()
	for i, line := range lines {
		if i < len(gradient) {
			fmt.Printf("%s%s%s\n", gradient[i], line, reset)
		} else {
			fmt.Printf("%s%s%s\n", gradient[len(gradient)-1], line, reset)
		}
	}
	if Version != "" {
		fmt.Printf("\033[1;38;5;203m                    %s%s\n", Version, reset)
	}
	fmt.Printf("       \033[38;5;203mVersatile Offline Communication Assistant%s\n", reset)
	fmt.Println()
}

