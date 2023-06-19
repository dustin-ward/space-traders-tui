package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type status int
const (
	intro status = iota
	accountSelect
	home
)

type MainModel struct {
	Status status
	
}

func (m MainModel) Init() tea.Cmd {
	return nil
}

func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "enter":
			m.Status = accountSelect
			return m, nil
		}
	}
	return m, nil
}

func (m MainModel) View() string {
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

func main() {
	p := tea.NewProgram(MainModel{}, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}
