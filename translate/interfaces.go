package translate

import "context"

type Backend interface {
	Translate(ctx context.Context, text, source, target string) (string, error)
}

type Language struct {
	Code string
	Name string
}

type LanguageProvider interface {
	List() []Language
}
