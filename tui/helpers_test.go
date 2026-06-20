package tui

import (
	"context"
	"testing"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/danterolle/loqi/translate"
)

func newTestModel(t *testing.T) Model {
	t.Helper()
	m := newModel(context.Background(), translate.NewMockBackend(), translate.NewStaticLanguages())
	mm, _ := m.Update(tea.WindowSizeMsg{Width: 80, Height: 24})
	return mm.(Model)
}
