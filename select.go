package main

import (
	"bytes"
	"fmt"

	"github.com/pkg/term"
)

var (
	up      = []byte{27, 91, 65}
	down    = []byte{27, 91, 66}
	sigquit = []byte{3}
	sigill  = []byte{4}
	tab     = []byte{9}
	enter   = []byte{13}
	esc     = []byte{27}
)

func cursorShow() {
	fmt.Print("\x1b[?25h")
}

func cursorHide() {
	fmt.Print("\x1b[?25l")
}

func cursorUp(n int) {
	fmt.Printf("\x1b[%dA", n)
}

func cursorMostLeft() {
	fmt.Printf("\x1b[%dG", 0)
}

func getch() []byte {
	t, _ := term.Open("/dev/tty")
	defer t.Close()

	term.RawMode(t)

	bytes := make([]byte, 3)
	numRead, err := t.Read(bytes)

	t.Restore()

	if err != nil {
		return nil
	} else {
		return bytes[0:numRead]
	}
}

func selectItem(sessions []Session) *Session {
	selected := 0

	cursorHide()
	defer cursorShow()

	for {
		for index, session := range sessions {
			var buffer bytes.Buffer

			if index == selected {
				buffer.WriteString("x ")
			} else {
				buffer.WriteString("  ")
			}

			buffer.WriteString(session.name)

			if index < len(sessions)-1 {
				buffer.WriteString("\n")
			}

			fmt.Print(buffer.String())
		}

		c := getch()

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
