package translate

import "context"

type Translator struct {
	Backend   Backend
	Languages LanguageProvider
}

func NewTranslator(backend Backend, langs LanguageProvider) *Translator {
	return &Translator{
		Backend:   backend,
		Languages: langs,
	}
}

func (t *Translator) Translate(ctx context.Context, text, source, target string) (string, error) {
	return t.Backend.Translate(ctx, text, source, target)
}
