package main

import "testing"

func sessionsEqual(a Session, b Session) bool {
	return a.name == b.name && a.windows == b.windows && a.attached == b.attached
}

func TestCompilePattern(t *testing.T) {
	_, err := compilePattern()

	if err != nil {
		t.Errorf("Compiling the pattern failed: %s", err)
	}
}

func TestParseSessions(t *testing.T) {
	fixture := []string{
		"fo1o(0): 1 windows (yee datez u know) [50x10]",
		"b_a-r$: 2 windows (some date) [100x20] (attached)",
		""}

	expected := []Session{
		Session{
			name:     "fo1o(0)",
			windows:  1,
			attached: false},
		Session{
			name:     "b_a-r$",
			windows:  2,
			attached: true}}

	r, _ := compilePattern()
	received := parseSessions(fixture, r)

	if len(received) != len(expected) {
		t.Errorf("Parsed sessions are incorrect, expecteded %d entries, received %d.",
			len(expected),
			len(received))
	} else {
		for i := range received {
			if !sessionsEqual(expected[i], received[i]) {
				t.Errorf("Sessions do not equal:\nExpected: %v\nReceived: %v\n",
					expected[i],
					received[i])
			}
		}
	}
}
