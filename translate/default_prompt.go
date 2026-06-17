package translate

import "fmt"

type defaultPrompt struct{}

func NewDefaultPrompt() *defaultPrompt {
	return &defaultPrompt{}
}

func (p *defaultPrompt) System() string {
	return "You are a translator. Translate the user's text accurately. Preserve meaning, tone, and sentence structure. Output only the translation — no greetings, explanations, or commentary."
}

func (p *defaultPrompt) Translate(text, source, target string) string {
	src := languages[source]
	if source == "auto" {
		src = "Auto detect"
	}
	return fmt.Sprintf("Translate from %s to %s:\n\n%s", src, languages[target], text)
}
