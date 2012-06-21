package desktopentry

import (
	"testing"
	"reflect"
)

func TestParseDesktopEntryEmpty(t *testing.T) {
	entry, err := ParseDesktopEntry("")
	if err != nil {
		t.Errorf("error: %s", err)
	}
	if !reflect.DeepEqual(entry, DesktopEntry{}) {
		t.Errorf("%v+ ≠ DesktopEntry{}", entry)
	}
}

/*
func TestParseDesktopEntryOneGroup(t *testing.T) {
	input := `[foo]
sense of life = 42
spam = eggs`
	expectedDesktopEntry := DesktopEntry{
		"foo": Group{
			Entry{"sense of life", "42"},
			Entry{"spam", "eggs"}}}
	entry, err := ParseDesktopEntry(input)
	if err != nil {
		t.Errorf("error: %s", err)
	}
	if !reflect.DeepEqual(entry, expectedDesktopEntry) {
		t.Errorf("%+v ≠ %+v", entry, expectedDesktopEntry)
	}
}
*/

/*
func TestParseDesktopEntryMultipleGroups(t *testing.T) {
	input := `[foobar]
x = 23
y = 42

[spam]
eggs = moo
ham = oink`
	expectedDesktopEntry := DesktopEntry{
		"foobar": Group{
			Entry{"x", "23"},
			Entry{"y", "42"}},
		"spam" : Group{
			Entry{"eggs", "moo"},
			Entry{"ham", "oink"}}}
	entry, err := ParseDesktopEntry(input)
	if err != nil {
		t.Errorf("error: %s", err)
	}
	if !reflect.DeepEqual(entry, expectedDesktopEntry) {
		t.Errorf("%+v ≠ %+v", entry, expectedDesktopEntry)
	}
}
*/

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
