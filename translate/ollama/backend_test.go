package ollama

import (
	"testing"

	"github.com/danterolle/voca/translate"
)

func TestBackend_ImplementsTranslateBackend(t *testing.T) {
	var _ translate.Backend = (*Backend)(nil)
}
