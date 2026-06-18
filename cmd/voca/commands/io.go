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

	stat, err := os.Stdin.Stat()
	if err != nil {
		return "", nil
	}
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		data, err := io.ReadAll(os.Stdin)
		if err != nil {
			return "", nil
		}
		return strings.TrimSpace(string(data)), nil
	}
	return "", nil
}

func ReadStdinOrFile(args []string) ([]byte, error) {
	if len(args) > 0 {
		path := args[0]
		if path == "-" {
			return io.ReadAll(os.Stdin)
		}
		return os.ReadFile(path)
	}

	stat, err := os.Stdin.Stat()
	if err != nil {
		return nil, fmt.Errorf("stdin not available: %w", err)
	}
	if (stat.Mode() & os.ModeCharDevice) != 0 {
		return nil, fmt.Errorf("no input file specified and stdin is a terminal")
	}
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		return nil, fmt.Errorf("read stdin: %w", err)
	}
	return data, nil
}
