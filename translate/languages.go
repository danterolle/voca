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

var (
	langCodes []string
	langNames map[string]string
)

func init() {
	codes := make([]string, 0, len(languages))
	for code := range languages {
		codes = append(codes, code)
	}
	sort.Strings(codes)
	langCodes = codes

	names := make(map[string]string, len(languages))
	for k, v := range languages {
		names[k] = v
	}
	langNames = names
}

type staticLanguages struct{}

func NewStaticLanguages() *staticLanguages {
	return &staticLanguages{}
}

func (s *staticLanguages) List() []Language {
	result := make([]Language, len(langCodes))
	for i, code := range langCodes {
		result[i] = Language{Code: code, Name: langNames[code]}
	}
	return result
}
