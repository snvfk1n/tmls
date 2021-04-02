package main

import (
	"bytes"
	"fmt"
	"strings"
)

func resetOutput(lines int) {
	cursorMostLeft()

	for i := 0; i < lines; i++ {
		eraseLine()

		if i < lines-1 {
			cursorUp(1)
		}
	}
}

func displaySessions(sessions []Session, selected int, length int) {
	for index, session := range sessions {
		var buffer bytes.Buffer

		if index == selected {
			buffer.WriteString("\x1b[7m")
		}

		label := []rune(session.Name)

		if len(label) > length {
			label = append(label[:length-1], '…')
		} else {
			space := strings.Repeat(" ", length-len(label))
			label = append(label, []rune(space)...)
		}

		buffer.WriteString(string(label))

		if index < len(sessions)-1 {
			buffer.WriteString("\n")
		}

		buffer.WriteString("\x1b[0m")

		eraseLine()
		fmt.Print(buffer.String())
	}
}

func selectItem(sessions []Session) *Session {
	if len(sessions) == 0 {
		fmt.Printf("There are no active tmux sessions.\n")
		return nil
	}

	selected := 0

	cursorHide()
	defer cursorShow()
	defer resetOutput(len(sessions))

	for {
		displaySessions(sessions, selected, 16)

		c := getCharacter()

		switch {
		case bytes.Equal(c, sigquit) || bytes.Equal(c, sigill) || bytes.Equal(c, quit):
			return nil
		case (bytes.Equal(c, up) || bytes.Equal(c, key_k)) && selected > 0:
			selected--
		case (bytes.Equal(c, down) || bytes.Equal(c, key_j)) && selected < len(sessions)-1:
			selected++
		case bytes.Equal(c, enter):
			return &sessions[selected]
		default:
		}

		cursorUp(len(sessions) - 1)
		cursorMostLeft()
	}

	return nil
}
