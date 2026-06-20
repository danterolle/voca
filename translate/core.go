package translate

import "context"

type Translator struct {
	backend   Backend
	languages LanguageProvider
}

func NewTranslator(backend Backend, langs LanguageProvider) *Translator {
	return &Translator{
		backend:   backend,
		languages: langs,
	}
}

func (t *Translator) Translate(ctx context.Context, text, source, target string) (string, error) {
	return t.backend.Translate(ctx, text, source, target)
}

func (t *Translator) Backend() Backend {
	return t.backend
}

func (t *Translator) Languages() LanguageProvider {
	return t.languages
}
