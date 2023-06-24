package auth

import tea "github.com/charmbracelet/bubbletea"

type (
	BackMsg struct{}
	ContMsg struct{}
)

type AuthModel struct{}

func NewAuthModel() *AuthModel {
	return &AuthModel{}
}

func (m AuthModel) Init() tea.Cmd {
	return nil
}

func (m AuthModel) View() string {
	return "LOGIN"
}

func (m AuthModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "esc":
			return m, func() tea.Msg { return BackMsg{} }
		case "enter":
			return m, func() tea.Msg { return ContMsg{} }
		}
	}
	return m, nil
}
