package translate

type Core struct {
	Backend   Backend
	Languages LanguageProvider
}

func NewCore(backend Backend, langs LanguageProvider) *Core {
	return &Core{
		Backend:   backend,
		Languages: langs,
	}
}
