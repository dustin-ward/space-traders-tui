package auth

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/dustin-ward/space-traders-tui/internal/constants"
)

var accountListStyle = lipgloss.NewStyle().Margin(1, 2)

type Account struct {
	Callsign string
	Faction  string
	Token    string
}

func (a Account) Title() string       { return a.Callsign }
func (a Account) Description() string { return a.Faction }
func (a Account) FilterValue() string { return a.Callsign }

type (
	BackMsg    struct{}
	ContMsg    struct{}
	SuccessMsg struct{ Account }
)

type authState int

const (
	selectState authState = iota
	createState
)

type AuthModel struct {
	state       authState
	form        tea.Model
	accountList list.Model
}

func NewAuthModel() *AuthModel {
	m := &AuthModel{
		state: selectState,
		accountList: list.New([]list.Item{
			Account{Callsign: "SHRUB", Faction: "VOID", Token: "st;oij3ltk3jtkdfjf3"},
			Account{Callsign: "TestCallsign1", Faction: "TEST", Token: "st;oij3ltk3jtkdfjf3"},
		}, list.NewDefaultDelegate(), 20, 30),
		form: NewCreateModel(),
	}
	if constants.WindowSize.Height != 0 {
		m.accountList.SetSize(constants.WindowSize.Width, constants.WindowSize.Height)
	}
	m.accountList.Title = "Choose Your Account"
	return m
}

func (m AuthModel) Init() tea.Cmd {
	return nil
}

func (m AuthModel) View() string {
	switch m.state {
	case selectState:
		return accountListStyle.Render(m.accountList.View())
	case createState:
		return m.form.View()
	}
	return "ERROR"
}

func (m AuthModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch m.state {
		case selectState:
			switch msg.String() {
			case "ctrl+c", "q":
				return m, tea.Quit
			case "esc":
				return m, func() tea.Msg { return BackMsg{} }
			case "enter":
				return m, func() tea.Msg { return ContMsg{} }
			case "c":
				m.state = createState
			}
		}
	case tea.WindowSizeMsg:
		h, v := accountListStyle.GetFrameSize()
		m.accountList.SetSize(msg.Width-h, msg.Height-v)
	case CancelCreateMsg:
		m.state = selectState
	case CreatedMsg:
		m.state = selectState
	}

	var cmd tea.Cmd
	switch m.state {
	case selectState:
		m.accountList, cmd = m.accountList.Update(msg)
		return m, cmd

	case createState:
		m.form, cmd = m.form.Update(msg)
		return m, cmd
	}
	return m, nil
}
