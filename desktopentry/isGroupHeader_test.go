package desktopentry

import "testing"

func TestIsGroupHeaderValid(t *testing.T) {
	line := "[the header]"
	if !isGroupHeader(line) {
		t.Errorf("%v is not an allowed group header", line)
	}
}

func TestIsGroupHeaderInvalid(t *testing.T) {
	lines := []string{
		"",
		"[]",
		"the header]",
		"[the header",
		"the header",
		"[[the header]]",
		"[the	header]"}
	for _, line := range lines {
		if isGroupHeader(line) {
			msg := "%v is an allowed group header, " +
				"although it shouldn't be"
			t.Errorf(msg, line)
		}
	}
}