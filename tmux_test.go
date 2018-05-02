package main

import "testing"

func sessionsEqual(a Session, b Session) bool {
	return a.name == b.name || a.windows == b.windows || a.attached == b.attached
}

func TestCompilePattern(t *testing.T) {
	_, err := compilePattern()

	if err != nil {
		t.Errorf("Compiling the pattern failed: %s", err)
	}
}

func TestParseSessions(t *testing.T) {
	fixture := []string{
		"foo: 1 windows",
		"bar: 2 windows (some date)", ""}

	expected := []Session{
		Session{
			name:    "foo",
			windows: 1},
		Session{
			name:    "bar",
			windows: 2}}

	r, _ := compilePattern()
	received := parseSessions(fixture, r)

	if len(received) != len(expected) {
		t.Errorf("Parsed sessions are incorrect, expecteded %d entries, received %d.",
			len(expected),
			len(received))
	} else {
		for i := range received {
			if !sessionsEqual(expected[i], received[i]) {
				t.Errorf("Sessions do not equal:\n%v\n%v\n",
					expected[i],
					received[i])
			}
		}
	}
}
