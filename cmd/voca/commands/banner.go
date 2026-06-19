package commands

import "fmt"

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
