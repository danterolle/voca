package commands

func printBanner() {
	if Quiet {
		return
	}

	gradient := []string{
		"\033[38;5;255m",
		"\033[38;5;230m",
		"\033[38;5;229m",
		"\033[38;5;227m",
		"\033[38;5;221m",
		"\033[38;5;215m",
		"\033[38;5;209m",
		"\033[38;5;203m",
	}
	reset := "\033[0m"

	lines := []string{
		"dP                   oo",
		"88                     ",
		"88 .d8888b. .d8888b. dP",
		"88 88'  `88 88'  `88 88",
		"88 88.  .88 88.  .88 88",
		"dP `88888P' `8888P88 dP",
		"                  88   ",
		"                  dP    ",
	}

	logDiag("\n")
	for i, line := range lines {
		if i < len(gradient) {
			logDiag("%s%s%s\n", gradient[i], line, reset)
		} else {
			logDiag("%s%s%s\n", gradient[len(gradient)-1], line, reset)
		}
	}
	logDiag("\n")
	if Version != "" {
		logDiag("\033[1;38;5;203m       %s%s\n", Version, reset)
	}
	logDiag("   \033[38;5;203mLOcal Quiet Interpreter%s\n", reset)
	logDiag("\n")
}
