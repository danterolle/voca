package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/danterolle/voca/translate"
)

var sentences = []string{
	"I would like to book a table for two at seven o'clock this evening, preferably near the window with a view of the garden.",
	"The conference will cover topics such as artificial intelligence, machine learning, and data privacy regulations across different countries.",
	"The museum's new exhibition features over two hundred paintings and sculptures from Renaissance artists, attracting visitors from all over the world.",
}

var targets = []struct {
	code string
	name string
}{
	{"it", "Italian"},
	{"fr", "French"},
	{"de", "German"},
	{"es", "Spanish"},
	{"pt", "Portuguese"},
	{"nl", "Dutch"},
	{"pl", "Polish"},
	{"ru", "Russian"},
	{"ja", "Japanese"},
	{"zh", "Chinese"},
	{"ko", "Korean"},
	{"ar", "Arabic"},
	{"tr", "Turkish"},
	{"hi", "Hindi"},
}

func main() {
	model := flag.String("model", translate.DefaultModel, "Ollama model")
	flag.Parse()

	totalStart := time.Now()
	totalSentences := len(sentences)
	langTimes := make(map[string][]time.Duration)

	for si, text := range sentences {
		fmt.Fprintf(os.Stderr, "\n=== Sentence %d/%d: %q ===\n", si+1, totalSentences, ellipsis(text, 60))
		for _, tgt := range targets {
			start := time.Now()
			result, err := translate.Translate(text, "en", tgt.code, *model)
			elapsed := time.Since(start)
			langTimes[tgt.code] = append(langTimes[tgt.code], elapsed)
			if err != nil {
				fmt.Printf("[%-4s %-10s] S%d ERROR: %v\n", tgt.code, tgt.name, si+1, err)
			} else {
				fmt.Printf("[%-4s %-10s] S%d %v — %q\n", tgt.code, tgt.name, si+1, elapsed.Round(time.Millisecond), result)
			}
		}
	}

	totalElapsed := time.Since(totalStart)
	fmt.Fprintf(os.Stderr, "\n=== Summary for %s ===\n", *model)
	fmt.Fprintf(os.Stderr, "Total time: %v | Avg per sentence: %v\n", totalElapsed.Round(time.Millisecond), (totalElapsed / time.Duration(totalSentences)).Round(time.Millisecond))
	fmt.Fprintf(os.Stderr, "\nLanguage averages:\n")
	for _, tgt := range targets {
		times := langTimes[tgt.code]
		if len(times) == 0 {
			continue
		}
		var sum time.Duration
		for _, t := range times {
			sum += t
		}
		avg := sum / time.Duration(len(times))
		fmt.Fprintf(os.Stderr, "  %-4s %-10s avg %v\n", tgt.code, tgt.name, avg.Round(time.Millisecond))
	}
}

func ellipsis(s string, max int) string {
	if len(s) <= max {
		return s
	}
	return s[:max-3] + "..."
}
