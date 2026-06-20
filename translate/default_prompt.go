package translate

import "fmt"

type chatPrompt struct{}

func NewChatPrompt() *chatPrompt {
	return &chatPrompt{}
}

func (p *chatPrompt) System() string {
	return "You are a translator. Translate the user's text accurately. Preserve meaning, tone, and sentence structure. Output only the translation — no greetings, explanations, or commentary."
}

func (p *chatPrompt) Translate(text, source, target string) string {
	src := languages[source]
	if source == "auto" {
		src = "Auto detect"
	}
	return fmt.Sprintf("Translate from %s to %s:\n\n%s", src, languages[target], text)
}
