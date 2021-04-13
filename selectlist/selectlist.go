package selectlist

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"strings"
)

type Model struct {
	Items  []string
	Cursor int
}

var subtle = lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#383838"}
var boxStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(subtle)
var activeStyle = lipgloss.NewStyle().Width(32).Background(subtle)

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
	items := make([]string, len(m.Items))
	for i, item := range m.Items {
		if i == m.Cursor {
			items[i] = activeStyle.Render(item)
		} else {
			items[i] = item
		}
	}
	v += boxStyle.Render(strings.Join(items, "\n"))
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
