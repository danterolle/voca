package main

import (
	"fmt"
	"time"

	"github.com/danterolle/voca/translate"
)

func main() {
	targets := []struct {
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

	text := "The quick brown fox jumps over the lazy dog. This sentence contains every letter of the English alphabet."

	for _, tgt := range targets {
		start := time.Now()
		result, err := translate.Translate(text, "en", tgt.code, translate.DefaultModel)
		elapsed := time.Since(start)
		if err != nil {
			fmt.Printf("[%-4s %-10s] ERROR: %v\n", tgt.code, tgt.name, err)
		} else {
			fmt.Printf("[%-4s %-10s] %v — %q\n", tgt.code, tgt.name, elapsed.Round(time.Millisecond), result)
		}
	}
}
