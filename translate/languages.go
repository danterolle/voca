package translate

import "sort"

var languages = map[string]string{
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

type staticLanguages struct {
	codes []string
	names map[string]string
}

func NewStaticLanguages() *staticLanguages {
	codes := make([]string, 0, len(languages))
	for code := range languages {
		codes = append(codes, code)
	}
	sort.Strings(codes)
	return &staticLanguages{
		codes: codes,
		names: languages,
	}
}

func (s *staticLanguages) List() []Language {
	result := make([]Language, len(s.codes))
	for i, code := range s.codes {
		result[i] = Language{Code: code, Name: s.names[code]}
	}
	return result
}
