package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func readInput(args []string) (string, error) {
	if len(args) > 0 {
		path := args[0]
		if info, err := os.Stat(path); err == nil && !info.IsDir() {
			data, err := os.ReadFile(path)
			if err != nil {
				return "", fmt.Errorf("reading file %q: %w", path, err)
			}
			return strings.TrimSpace(string(data)), nil
		}
		return strings.Join(args, " "), nil
	}
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		data, err := io.ReadAll(os.Stdin)
		if err != nil {
			return "", fmt.Errorf("reading stdin: %w", err)
		}
		return strings.TrimSpace(string(data)), nil
	}
	return "", nil
}

func readStdinOrFile(args []string) ([]byte, error) {
	if len(args) > 0 {
		return os.ReadFile(args[0])
	}
	return io.ReadAll(os.Stdin)
}
