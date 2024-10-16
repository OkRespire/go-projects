package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	query  []string
	width  int
	height int
}

func New(query []string) *Model {
	return &Model{query: query}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c": // allows the program to quit.
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m Model) View() string {
	if m.width == 0 {
		return "loading"
	}

	return "loaded"
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
		panic(err)
	}
}
