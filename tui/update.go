package tui

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		return m.handleWindowSize(msg)

	case tea.KeyMsg:
		return m.handleKey(msg)

	case debounceMsg:
		return m.handleDebounce(msg)

	case translateResultMsg:
		return m.handleTranslateResult(msg)
	}

	var cmd tea.Cmd
	m.textarea, cmd = m.textarea.Update(msg)
	return m, cmd
}

func (m Model) handleWindowSize(msg tea.WindowSizeMsg) (Model, tea.Cmd) {
	m.ready = true
	m.width = msg.Width
	contentH := msg.Height - 12
	if contentH < 4 {
		contentH = 4
	}
	m.textarea.SetWidth(msg.Width - 6)
	m.textarea.SetHeight(contentH / 2)
	return m, nil
}

func (m Model) handleDebounce(msg debounceMsg) (Model, tea.Cmd) {
	if msg.seq != m.translateSequence {
		return m, nil
	}
	m.leadingInProgress = false
	text := m.textarea.Value()
	if text == "" || text == m.lastInput {
		return m, nil
	}
	m.lastInput = text
	sourceCode := m.langCodes[m.sourceIdx]
	targetCode := m.langCodes[m.targetIdx]
	m.status = fmt.Sprintf("Translating... (%s -> %s)", m.langNames[sourceCode], m.langNames[targetCode])
	return m, m.doTranslate(text, sourceCode, targetCode)
}

func (m Model) handleTranslateResult(msg translateResultMsg) (Model, tea.Cmd) {
	if msg.text != m.textarea.Value() {
		return m, nil
	}
	if msg.err != nil {
		m.status = fmt.Sprintf("Error: %v", msg.err)
	} else {
		if msg.result != "" {
			m.output = msg.result
		}
		m.status = "Ready."
	}
	return m, nil
}

func (m Model) handleKey(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	if m.focused == focusSrcLang || m.focused == focusTgtLang {
		return m.handleLangKey(msg)
	}
	return m.handleInputKey(msg)
}

func (m Model) handleLangKey(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "tab":
		return m.advanceFocus(), nil
	case "shift+tab":
		return m.retreatFocus(), nil
	case "up":
		m = m.adjustLangIndex(-1)
		return m, m.translateOnLangChange()
	case "down":
		m = m.adjustLangIndex(1)
		return m, m.translateOnLangChange()
	case "ctrl+c", "esc":
		return m, tea.Quit
	}
	return m, nil
}

func (m Model) adjustLangIndex(delta int) Model {
	switch m.focused {
	case focusSrcLang:
		idx := m.sourceIdx + delta
		if idx >= 0 && idx < len(m.langCodes) {
			m.sourceIdx = idx
		}
	case focusTgtLang:
		idx := m.targetIdx + delta
		if idx >= 0 && idx < len(m.langCodes) {
			m.targetIdx = idx
		}
	}
	return m
}

func (m Model) handleInputKey(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "tab":
		return m.advanceFocus(), nil
	case "shift+tab":
		return m.retreatFocus(), nil
	case "ctrl+y":
		if m.output != "" {
			if err := copyClipboard(m.output); err != nil {
				m.status = fmt.Sprintf("Clipboard error: %v", err)
			} else {
				m.status = "Copied to clipboard!"
			}
		}
		return m, nil
	case "ctrl+l":
		m.textarea.Reset()
		m.output = ""
		m.lastInput = ""
		m.translateSequence++
		m.leadingInProgress = false
		m.status = "Cleared."
		return m, nil
	case "ctrl+t":
		if m.langCodes[m.sourceIdx] != "auto" {
			m.sourceIdx, m.targetIdx = m.targetIdx, m.sourceIdx
			m.status = "Languages swapped."
			return m, m.translateOnLangChange()
		}
		return m, nil
	case "ctrl+c", "esc":
		return m, tea.Quit
	}

	return m.handleTextChange(msg)
}

func (m Model) handleTextChange(msg tea.KeyMsg) (Model, tea.Cmd) {
	var cmd tea.Cmd
	before := m.textarea.Value()
	m.textarea, cmd = m.textarea.Update(msg)
	if m.textarea.Value() != before {
		m.translateSequence++
		if m.leadingInProgress {
			cmd = tea.Batch(cmd, m.scheduleDebounce(m.translateSequence))
		} else {
			m, cmd = m.startLeadingTranslate(cmd)
		}
	}
	return m, cmd
}

func (m Model) scheduleDebounce(seq int) tea.Cmd {
	return tea.Tick(debounceDuration, func(t time.Time) tea.Msg {
		return debounceMsg{seq: seq}
	})
}

func (m Model) translateOnLangChange() tea.Cmd {
	text := m.textarea.Value()
	if text == "" {
		return nil
	}
	sourceCode := m.langCodes[m.sourceIdx]
	targetCode := m.langCodes[m.targetIdx]
	m.status = fmt.Sprintf("Translating... (%s -> %s)", m.langNames[sourceCode], m.langNames[targetCode])
	return m.doTranslate(text, sourceCode, targetCode)
}

func (m Model) startLeadingTranslate(prevCmd tea.Cmd) (Model, tea.Cmd) {
	text := m.textarea.Value()
	if text == "" {
		return m, prevCmd
	}
	m.leadingInProgress = true
	m.lastInput = text
	sourceCode := m.langCodes[m.sourceIdx]
	targetCode := m.langCodes[m.targetIdx]
	m.status = fmt.Sprintf("Translating... (%s -> %s)", m.langNames[sourceCode], m.langNames[targetCode])
	return m, tea.Batch(prevCmd, m.doTranslate(text, sourceCode, targetCode))
}

func (m Model) advanceFocus() Model {
	m.textarea.Blur()
	switch m.focused {
	case focusSrcLang:
		m.focused = focusTgtLang
	case focusTgtLang:
		m.focused = focusInput
		m.textarea.Focus()
	case focusInput:
		m.focused = focusSrcLang
	}
	return m
}

func (m Model) retreatFocus() Model {
	m.textarea.Blur()
	switch m.focused {
	case focusSrcLang:
		m.focused = focusInput
		m.textarea.Focus()
	case focusTgtLang:
		m.focused = focusSrcLang
	case focusInput:
		m.focused = focusTgtLang
	}
	return m
}
