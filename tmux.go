package main

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

// Session describes a tmux session.
type Session struct {
	name     string
	windows  int
	attached bool
}

func compilePattern() (*regexp.Regexp, error) {
	return regexp.Compile(`^(\w+)\: (\d+) windows`)
}

func getTmuxSessions() []string {
	var (
		output []byte
		err    error
	)

	if output, err = exec.Command("tmux", []string{"ls"}...).Output(); err != nil {
		fmt.Fprintln(os.Stderr, "There was an error running `tmux ls`: ", err)
		os.Exit(1)
	}

	return strings.Split(string(output), "\n")
}

func parseSessions(sessionEntries []string, r *regexp.Regexp) []Session {
	var (
		sessions []Session
	)

	for _, sessionEntry := range sessionEntries {
		res := r.FindAllStringSubmatch(sessionEntry, -1)

		if len(res) != 1 {
			fmt.Printf("heh wat? (1) (%d results)\n", len(res))
			break
		}

		match := res[0]

		if len(match) != 3 {
			fmt.Println("heh wat? (2)")
			continue
		}

		windows, err := strconv.Atoi(match[2])

		if err != nil {
			fmt.Println("heh wat? (3)")
			continue
		}

		sessions = append(sessions, Session{
			name:     match[1],
			windows:  windows,
			attached: false})
	}

	return sessions
}

func getSessions() []Session {
	r, err := compilePattern()

	if err != nil {
		fmt.Fprintln(os.Stderr, "There is an error with our expression: ", err)
		os.Exit(1)
	}

	sessionEntries := getTmuxSessions()
	return parseSessions(sessionEntries, r)
}
