package main

import (
	"fmt"
	input "github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/jankaszel/tmls/selectlist"
	"log"
	"os"
	"path/filepath"
	"slices"
)

const (
	ModeSelect = "SELECT_SESSION"
	ModeCreate = "CREATE_SESSION"
)

type model struct {
	sessions   []Session
	selectList selectlist.Model
	nameInput  input.Model
	mode       string

	err error
}
type errMsg error

func (m model) Init() tea.Cmd {
	return input.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit

		case tea.KeyEnter:
			if m.mode == ModeCreate {
				sessionName := m.nameInput.Value()
				createSession(sessionName)
			} else if m.mode == ModeSelect {
				session := m.sessions[m.selectList.Cursor]
				attachSession(&session)
			}
			return m, tea.Quit

		case tea.KeyTab:
			if m.mode == ModeCreate {
				m.mode = ModeSelect
				m.nameInput.Blur()
				return m, nil
			} else {
				m.mode = ModeCreate
				m.nameInput.Focus()
				return m, cmd
			}

		case tea.KeyRunes:
			if len(msg.Runes) == 1 {
				if msg.Runes[0] == 'q' && m.mode == ModeSelect {
					return m, tea.Quit
				}
			}
		}

	case errMsg:
		m.err = msg
		return m, nil
	}

	if m.mode == ModeCreate {
		m.nameInput, cmd = m.nameInput.Update(msg)
	} else if m.mode == ModeSelect {
		m.selectList, cmd = m.selectList.Update(msg)
	}

	return m, cmd
}

func (m model) View() string {
	return fmt.Sprintf(
		"%s\n%s\n",
		m.selectList.View(),
		m.nameInput.View(),
	)
}

func sessionNames(sessions []Session) []string {
	names := make([]string, len(sessions))
	for i, session := range sessions {
		names[i] = session.Name
	}
	return names
}

func initialModel() model {
	sessions := getSessions()

	nameInput := input.NewModel()
	nameInput.Placeholder = "Create session"
	nameInput.PlaceholderColor = "237"
	nameInput.CharLimit = 64
	nameInput.Width = 30

	path, err := os.Getwd()
	if err == nil {
		base := filepath.Base(path)
		idx := slices.IndexFunc(sessions, func(s Session) bool { return s.Name == base })

		if idx == -1 {
			nameInput.SetValue(base)
			nameInput.CursorEnd()
			nameInput.Blur()
		}
	}

	mode := ModeSelect
	selectList := selectlist.NewModel(sessionNames(sessions))

	if len(sessions) == 0 {
		mode = ModeCreate
		nameInput.Focus()
	}

	return model{
		sessions:   sessions,
		selectList: selectList,
		nameInput:  nameInput,
		mode:       mode,
		err:        nil,
	}
}

func main() {
	p := tea.NewProgram(initialModel())

	if err := p.Start(); err != nil {
		log.Fatal(err)
	}
}
