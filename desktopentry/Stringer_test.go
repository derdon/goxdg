package desktopentry

import (
	"testing"
	"fmt"
)

func TestGroupDeclarationOnly(t *testing.T) {
	entries := Entries{"meh": Group{}}
	output := fmt.Sprintf("%s", entries)
	expectedOutput := "[meh]"
	if output != expectedOutput {
		t.Errorf("'%s' != '%s'", output, expectedOutput)
	}
}

func TestWithOneEntry(t *testing.T) {
	entries := Entries{"foo": Group{Entry{"sense of life", "42"}}}
	output := fmt.Sprintf("%s", entries)
	expectedOutput := `[foo]
sense of life = 42`
	if output != expectedOutput {
		t.Errorf("'%s' != '%s'", output, expectedOutput)
	}	
}

func TestWithMultipleEntries(t *testing.T) {
	entries := Entries{"bar": Group{Entry{"name", "bob"}, Entry{"age", "23"}}}
	output := fmt.Sprintf("%s", entries)
	expectedOutput := `[bar]
name = bob
age = 23`
	if output != expectedOutput {
		t.Errorf("'%s' != '%s'", output, expectedOutput)
	}	
}

func TestWithMultipleGroups(t *testing.T) {
	entries := Entries{
		"group1": Group{Entry{"foo", "bar"}},
		"group2": Group{Entry{"spam", "eggs"}}}
	output := fmt.Sprintf("%s", entries)
	possibleOutput1 := `[group1]
foo = bar
[group2]
spam = eggs`
	possibleOutput2 := `[group2]
spam = eggs
[group1]
foo = bar`
	if output != possibleOutput1 && output != possibleOutput2 {
		t.Errorf(
			"'%s' != '%s' && %s != %s",
			output, possibleOutput1, output, possibleOutput2)
	}	
}