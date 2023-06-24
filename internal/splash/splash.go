package splash

import (
	tea "github.com/charmbracelet/bubbletea"
)

type (
	ContMsg struct{}
)

type SplashModel struct{}

func NewSplashModel() *SplashModel {
	return &SplashModel{}
}

func (m SplashModel) Init() tea.Cmd {
	return nil
}

func (m SplashModel) View() string {
	s := `
░██████╗██████╗░░█████╗░░█████╗░███████╗  ████████╗██████╗░░█████╗░██████╗░███████╗██████╗░░██████╗
██╔════╝██╔══██╗██╔══██╗██╔══██╗██╔════╝  ╚══██╔══╝██╔══██╗██╔══██╗██╔══██╗██╔════╝██╔══██╗██╔════╝
╚█████╗░██████╔╝███████║██║░░╚═╝█████╗░░  ░░░██║░░░██████╔╝███████║██║░░██║█████╗░░██████╔╝╚█████╗░
░╚═══██╗██╔═══╝░██╔══██║██║░░██╗██╔══╝░░  ░░░██║░░░██╔══██╗██╔══██║██║░░██║██╔══╝░░██╔══██╗░╚═══██╗
██████╔╝██║░░░░░██║░░██║╚█████╔╝███████╗  ░░░██║░░░██║░░██║██║░░██║██████╔╝███████╗██║░░██║██████╔╝
╚═════╝░╚═╝░░░░░╚═╝░░╚═╝░╚════╝░╚══════╝  ░░░╚═╝░░░╚═╝░░╚═╝╚═╝░░╚═╝╚═════╝░╚══════╝╚═╝░░╚═╝╚═════╝░
`
	s += "\n\n"
	s += "Press enter to continue\n"
	s += "Press q to quit...\n"
	return s
}

func (m SplashModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "enter":
			return m, func() tea.Msg { return ContMsg{} }
		}
	}
	return m, nil
}
