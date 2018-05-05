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
	return regexp.Compile(`^(\w+)\: (\d+) windows \(.*\) \[(\d+)x(\d+)\](\s\(attached\))?`)
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
			continue
		}

		match := res[0]

		if len(match) < 5 {
			continue
		}

		windows, err := strconv.Atoi(match[2])

		if err != nil {
			continue
		}

		sessions = append(sessions, Session{
			name:     match[1],
			windows:  windows,
			attached: len(match) > 5 && match[5] != ""})
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

func attachSession(session *Session) {
	fmt.Println(session.name)

	cmd := exec.Command("tmux", "attach", "-t", session.name)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.Run()
}
