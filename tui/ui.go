package tui

import (
	"context"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/danterolle/voca/translate"
)

type UI interface {
	Run(ctx context.Context, core *translate.Core) error
}

type BubbleTeaUI struct{}

func NewBubbleTeaUI() *BubbleTeaUI {
	return &BubbleTeaUI{}
}

func (u *BubbleTeaUI) Run(ctx context.Context, core *translate.Core) error {
	m := newModel(core)
	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		return err
	}
	return nil
}
