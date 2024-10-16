package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/charmbracelet/bubbles/textinput"
)

type Styles struct {
	BorderCol  lipgloss.Color
	InputField lipgloss.Style
}

func DefaultStyles() *Styles {
	s := new(Styles)

	s.BorderCol = lipgloss.Color("57")

	s.InputField = lipgloss.NewStyle().BorderForeground(s.BorderCol).BorderStyle(lipgloss.NormalBorder()).Padding(1).Width(80)

	return s
}

type Model struct {
	query  []string
	width  int
	height int
	field  textinput.Model
	index  int
	styles *Styles
}

func New(query []string) *Model {
	field := textinput.New()
	field.Placeholder = "Query here"
	field.Focus()
	styles := DefaultStyles()
	return &Model{query: query, field: field, styles: styles}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c": // allows the program to quit.
			return m, tea.Quit
		case "enter":
			m.index++
			m.field.SetValue("done")
			return m, nil
		}

	}
	m.field, cmd = m.field.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	if m.width == 0 {
		return "loading"
	}

	return lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Center,
		lipgloss.Center,
		lipgloss.JoinVertical(
			lipgloss.Center,
			m.query[m.index],
			m.styles.InputField.Render(m.field.View()),
		),
	)
}

func main() {
	query := []string{"what is your name?"}
	m := New(query)

	f, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	p := tea.NewProgram(m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
