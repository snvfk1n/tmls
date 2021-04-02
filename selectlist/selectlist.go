package selectlist

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/muesli/termenv"
)

type Model struct {
	Items  []string
	Cursor int
}

var color func(s string) termenv.Color = termenv.ColorProfile().Color

func NewModel(items []string) Model {
	return Model{
		Items:  items,
		Cursor: 0,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) View() string {
	if len(m.Items) == 0 {
		return "No active sessions."
	}

	v := "Active sessions:\n"
	for i, item := range m.Items {
		if i == m.Cursor {
			v += item
		} else {
			v += termenv.String(item).Foreground(color("237")).String()
		}
		v += "\n"
	}
	return v
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "k":
			if m.Cursor > 0 {
				m.Cursor = m.Cursor - 1
			}

		case "down", "j":
			if m.Cursor < len(m.Items)-1 {
				m.Cursor = m.Cursor + 1
			}
		}
	}
	return m, nil
}
