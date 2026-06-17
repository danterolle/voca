package translate

import "sort"

const DefaultModel = "gemma4:e2b-it-qat"

var Languages = map[string]string{
	"auto": "Auto",
	"en":   "English",
	"it":   "Italian",
	"fr":   "French",
	"de":   "German",
	"es":   "Spanish",
	"pt":   "Portuguese",
	"nl":   "Dutch",
	"pl":   "Polish",
	"ru":   "Russian",
	"ja":   "Japanese",
	"zh":   "Chinese",
	"ko":   "Korean",
	"ar":   "Arabic",
	"tr":   "Turkish",
	"cs":   "Czech",
	"sv":   "Swedish",
	"da":   "Danish",
	"fi":   "Finnish",
	"el":   "Greek",
	"ro":   "Romanian",
	"hu":   "Hungarian",
	"vi":   "Vietnamese",
	"th":   "Thai",
	"hi":   "Hindi",
}

type StaticLanguages struct {
	codes []string
	names map[string]string
}

func NewStaticLanguages() *StaticLanguages {
	codes := make([]string, 0, len(Languages))
	for code := range Languages {
		codes = append(codes, code)
	}
	sort.Strings(codes)
	return &StaticLanguages{
		codes: codes,
		names: Languages,
	}
}

func (s *StaticLanguages) List() []Language {
	result := make([]Language, len(s.codes))
	for i, code := range s.codes {
		result[i] = Language{Code: code, Name: s.names[code]}
	}
	return result
}

func (s *StaticLanguages) Lookup(code string) (Language, bool) {
	name, ok := s.names[code]
	if !ok {
		return Language{}, false
	}
	return Language{Code: code, Name: name}, true
}
