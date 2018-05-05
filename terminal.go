package main

import (
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

func eraseLine() {
	fmt.Print("\x1b[2K")
}

func getCharacter() []byte {
	t, _ := term.Open("/dev/tty")
	defer t.Close()

	term.RawMode(t)

	bytes := make([]byte, 3)
	numRead, err := t.Read(bytes)

	t.Restore()

	if err != nil {
		return nil
	}

	return bytes[0:numRead]
}
