package main

import (
	tea "github.com/charmbracelet/bubbletea"

	"github.com/dustin-ward/space-traders-tui/internal/auth"
	"github.com/dustin-ward/space-traders-tui/internal/constants"
	"github.com/dustin-ward/space-traders-tui/internal/home"
	"github.com/dustin-ward/space-traders-tui/internal/splash"
)

type sessionState int

const (
	splashScreen sessionState = iota
	authScreen
	homeScreen
)

type MainModel struct {
	state  sessionState
	models map[sessionState]tea.Model
}

func (m MainModel) Init() tea.Cmd {
	return nil
}

func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		constants.WindowSize = msg
	case splash.ContMsg:
		m.state = authScreen
	case auth.ContMsg:
		m.state = homeScreen
	case auth.BackMsg:
		m.state = splashScreen
	case home.BackMsg:
		m.state = splashScreen
	}

	var cmd tea.Cmd
	switch m.state {
	case splashScreen:
		m.models[splashScreen], cmd = m.models[splashScreen].Update(msg)
		return m, cmd

	case authScreen:
		// m.models[authScreen] = *auth.NewAuthModel()
		m.models[authScreen], cmd = m.models[authScreen].Update(msg)
		return m, cmd

	case homeScreen:
		m.models[homeScreen], cmd = m.models[homeScreen].Update(msg)
		return m, cmd

	}
	return m, nil
}

func (m MainModel) View() string {
	return m.models[m.state].View()
}

func NewMainModel() *MainModel {
	return &MainModel{
		state: splashScreen,
		models: map[sessionState]tea.Model{
			splashScreen: splash.NewSplashModel(),
			authScreen:   auth.NewAuthModel(),
			homeScreen:   home.NewHomeModel(),
		},
	}
}
