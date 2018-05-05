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

		label := []rune(session.name)

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
	selected := 0

	cursorHide()
	defer cursorShow()
	defer resetOutput(len(sessions))

	for {
		displaySessions(sessions, selected, 16)

		c := getCharacter()

		switch {
		case bytes.Equal(c, sigquit) || bytes.Equal(c, sigill):
			return nil
		case bytes.Equal(c, up) && selected > 0:
			selected--
		case bytes.Equal(c, down) && selected < len(sessions)-1:
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
