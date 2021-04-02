package main

import "testing"

func sessionsEqual(a Session, b Session) bool {
	return a.Name == b.Name && a.Windows == b.Windows && a.Attached == b.Attached
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
			Name:     "fo1o(0)",
			Windows:  1,
			Attached: false},
		Session{
			Name:     "b_a-r$",
			Windows:  2,
			Attached: true}}

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
