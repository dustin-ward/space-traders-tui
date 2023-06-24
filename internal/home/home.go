package home

import tea "github.com/charmbracelet/bubbletea"

type (
	BackMsg struct{}
)

type HomeModel struct{}

func NewHomeModel() HomeModel {
	return HomeModel{}
}

func (m HomeModel) Init() tea.Cmd {
	return nil
}

func (m HomeModel) View() string {
	return "HOME"
}

func (m HomeModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "esc":
			return m, func() tea.Msg { return BackMsg{} }
		}
	}
	return m, nil
}
