package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"syscall"
)

// Session describes a tmux session.
type Session struct {
	Name     string
	Windows  int
	Attached bool
}

func compilePattern() (*regexp.Regexp, error) {
	return regexp.Compile(`^([^:]+)\: (\d+) windows \([^\(]*\)(\s\[(\d+)x(\d+)\])?(\s\(attached\))?`)
}

func getTmuxSessions() []string {
	var (
		output []byte
		err    error
	)

	if output, err = exec.Command("tmux", []string{"ls"}...).Output(); err != nil {
		if ee, ok := err.(*exec.ExitError); ok &&
			(strings.HasPrefix(string(ee.Stderr), "no server running on") ||
				strings.HasPrefix(string(ee.Stderr), "error connecting to")) {
			return []string{}
		}

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

		if len(match) < 3 {
			continue
		}

		windows, err := strconv.Atoi(match[2])
		attachedValue := match[len(match)-1]
		attached := attachedValue == " (attached)"

		if err != nil {
			continue
		}

		sessions = append(sessions, Session{
			Name:     match[1],
			Windows:  windows,
			Attached: attached})
	}

	return sessions
}

func getSessions() []Session {
	r, err := compilePattern()

	if err != nil {
		fmt.Fprintln(os.Stderr, "There is an error with our regular expression: ", err)
		os.Exit(1)
	}

	sessionEntries := getTmuxSessions()
	return parseSessions(sessionEntries, r)
}

func attachSession(session *Session) {
	fmt.Println(session.Name)

	tmuxPath, err := exec.LookPath("tmux")
	if err != nil {
		log.Fatalf("Could not find tmux: %v", err)
	}

	args := []string{"tmux", "-u", "attach", "-t", session.Name}

	if err := syscall.Exec(tmuxPath, args, os.Environ()); err != nil {
		log.Fatalf("Failed to exec tmux: %v", err)
	}
}

func createSession(name string) {
	tmuxPath, err := exec.LookPath("tmux")
	if err != nil {
		log.Fatalf("Could not find tmux: %v", err)
	}

	args := []string{"tmux", "-u", "new", "-s", name}

	if err := syscall.Exec(tmuxPath, args, os.Environ()); err != nil {
		log.Fatalf("Failed to exec tmux: %v", err)
	}
}
