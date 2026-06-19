package ollama

import (
	"fmt"
	"os"
	"strings"
)

func renderPullStatus(status string, total, completed int64) {
	if total > 0 {
		pct := float64(completed) / float64(total) * 100
		bar := progressBar(pct, 30)
		fmt.Fprintf(os.Stderr, "\r     %s  %.0f%%", bar, pct)
	} else if status == "success" {
		fmt.Fprintf(os.Stderr, "\r     %s  100%%\n", progressBar(100, 30))
	} else if strings.Contains(status, "pulling") {
		parts := strings.SplitN(status, " ", 2)
		if len(parts) == 2 {
			short := parts[1]
			if len(short) > 12 {
				short = short[:12]
			}
			fmt.Fprintf(os.Stderr, "\r     Pulling %s...", short)
		}
	} else if status == "verifying sha256 digest" {
		fmt.Fprintf(os.Stderr, "\r     Verifying...")
	} else if status == "writing manifest" {
		fmt.Fprintf(os.Stderr, "\r     Writing manifest...")
	} else {
		fmt.Fprintf(os.Stderr, "\r     %s", status)
	}
}

func progressBar(pct float64, width int) string {
	filled := int(pct * float64(width) / 100)
	if filled > width {
		filled = width
	}
	return strings.Repeat("█", filled) + strings.Repeat("░", width-filled)
}
