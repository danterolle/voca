package translate

import "fmt"

type DefaultPrompt struct{}

func NewDefaultPrompt() *DefaultPrompt {
	return &DefaultPrompt{}
}

func (p *DefaultPrompt) System() string {
	return "You are a translator. Translate the user's text accurately. Preserve meaning, tone, and sentence structure. Output only the translation — no greetings, explanations, or commentary."
}

func (p *DefaultPrompt) Translate(text, source, target string) string {
	src := Languages[source]
	if source == "auto" {
		src = "Auto detect"
	}
	return fmt.Sprintf("Translate from %s to %s:\n\n%s", src, Languages[target], text)
}
