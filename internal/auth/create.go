package auth

import (
	"fmt"

	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type (
	CancelCreateMsg struct{}
	CreatedMsg      struct{ Account }
)

type CreateModel struct {
	focusIdx   int
	input      textinput.Model
	list       list.Model
	cursorMode cursor.Mode
}

func NewCreateModel() *CreateModel {
	m := &CreateModel{
		input: textinput.New(),
	}

	m.input.Focus()
	m.input.CharLimit = 32
	m.input.Width = 20

	return m
}

func (m CreateModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m CreateModel) View() string {
	return fmt.Sprintf(
		"Create an Account\n\nCallsign:\n%s",
		m.input.View(),
	)
}

func (m CreateModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "enter":
			return m, func() tea.Msg { return ContMsg{} }
		case "esc":
			return m, func() tea.Msg { return CancelCreateMsg{} }
		}
	}

	var cmd tea.Cmd
	m.input, cmd = m.input.Update(msg)
	return m, cmd
}
