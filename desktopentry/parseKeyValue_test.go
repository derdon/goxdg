package desktopentry

import "testing"

func TestParseKeyValue(t *testing.T) {
	lines := []string{"foo=42", "foo  =       42"}
	for _, line := range lines {
		entry, err := parseKeyValue(line)
		if err != nil {
			t.Errorf("error: %s", err)
		}
		expectedEntry := Entry{"foo", "42"}
		if entry != expectedEntry {
			t.Errorf("%+v ≠ %+v", entry, expectedEntry)
		}
	}
}

func TestParseKeyValueMultiWordKey(t *testing.T) {
	entry, err := parseKeyValue("sense of life = 42")
	if err != nil {
		t.Errorf("error: %s", err)
	}
	expectedEntry := Entry{"sense of life", "42"}
	if entry != expectedEntry {
		t.Errorf("%+v ≠ %+v", entry, expectedEntry)
	}
}

func TestParseKeyValueInvalid(t *testing.T) {
	lines := []string{
		"bla bla",
		"    	",
		"",
		"=",
		"foo=",
		"=bar"}
	for _, line := range lines {
		entry, err := parseKeyValue(line)
		if (entry != Entry{}) {
			t.Errorf("%v ≠ Entry{}", entry)
		}
		if err == nil {
			t.Errorf("function should have ended with an error")
		}
	}
}

func TestParseKeyValueMultipleAssignOps(t *testing.T) {
	line := "a = b = c = d"
	entry, err := parseKeyValue(line)
	if err != nil {
		t.Errorf("error: %s", err)
	}
	expectedEntry := Entry{"a", "b = c = d"}
	if entry != expectedEntry {
		t.Errorf("%+v ≠ %+v", entry, expectedEntry)
	}
}