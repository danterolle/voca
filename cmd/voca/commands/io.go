package commands

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func ReadInput(args []string) (string, error) {
	if len(args) > 0 {
		data, err := ReadStdinOrFile(args)
		if err != nil {
			return "", err
		}
		return strings.TrimSpace(string(data)), nil
	}

	data, err := readStdin()
	if err != nil || data == nil {
		return "", nil
	}
	return strings.TrimSpace(string(data)), nil
}

func ReadStdinOrFile(args []string) ([]byte, error) {
	if len(args) > 0 {
		path := args[0]
		if path == "-" {
			return io.ReadAll(os.Stdin)
		}
		return os.ReadFile(path)
	}

	data, err := readStdin()
	if err != nil {
		return nil, fmt.Errorf("stdin not available: %w", err)
	}
	if data == nil {
		return nil, fmt.Errorf("no input file specified and stdin is a terminal")
	}
	return data, nil
}

func readStdin() ([]byte, error) {
	stat, err := os.Stdin.Stat()
	if err != nil {
		return nil, err
	}
	if (stat.Mode() & os.ModeCharDevice) != 0 {
		return nil, nil // terminal, no piped data
	}
	return io.ReadAll(os.Stdin)
}
