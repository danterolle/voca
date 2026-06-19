package commands

import (
	"fmt"
	"os"
)

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
		"__      ______   _____          ",
		"\\ \\    / / __ \\ / ____|   /\\    ",
		" \\ \\  / / |  | | |       /  \\   ",
		"  \\ \\/ /| |  | | |      / /\\ \\  ",
		"   \\  / | |__| | |____ / ____ \\ ",
		"    \\/   \\____/ \\_____/_/    \\_\\",
	}

	fmt.Fprintln(os.Stderr)
	for i, line := range lines {
		if i < len(gradient) {
			fmt.Fprintf(os.Stderr, "%s%s%s\n", gradient[i], line, reset)
		} else {
			fmt.Fprintf(os.Stderr, "%s%s%s\n", gradient[len(gradient)-1], line, reset)
		}
	}
	if Version != "" {
		fmt.Fprintf(os.Stderr, "\033[1;38;5;203m                    %s%s\n", Version, reset)
	}
	fmt.Fprintf(os.Stderr, "       \033[38;5;203mVersatile Offline Communication Assistant%s\n", reset)
	fmt.Fprintln(os.Stderr)
}
