package desktopentry

import (
	"testing"
	"reflect"
	"fmt"
)

func TestEmpty(t *testing.T) {
	entries, err := ParseDesktopEntryString("")
	if err != nil {
		t.Errorf("error: %s", err)
	}
	if !reflect.DeepEqual(entries, Entries{}) {
		t.Errorf("%v+ != Entries{}", entries)
	}
}

func TestParseGroupDeclarationOnly(t *testing.T) {
	entries, err := ParseDesktopEntryString("[Desktop Entry]")
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	expectedEntries := Entries{"Desktop Entry": Group{}}
	if !reflect.DeepEqual(entries, expectedEntries) {
		t.Errorf("%v != %v", entries, expectedEntries)
	}
}

func TestMissingGroupDeclaration(t *testing.T) {
	_, err := ParseDesktopEntryString("foo = 42")
	if err == nil {
		t.Error("expected to get an error, but got none")
	}
	errmsg := fmt.Sprintf("%s", err)
	expectedErrMsg := "missing group header before entry declarations (line 1)"
	if errmsg != expectedErrMsg {
		t.Errorf("%v != %v", errmsg, expectedErrMsg)
	}
}

func TestOneGroup(t *testing.T) {
	input := `[foo]
sense of life = 42
spam = eggs`
	expectedEntries := Entries{
		"foo": Group{
			Entry{"sense of life", "42"},
			Entry{"spam", "eggs"}}}
	entries, err := ParseDesktopEntryString(input)
	if err != nil {
		t.Errorf("error: %s", err)
	}
	if !reflect.DeepEqual(entries, expectedEntries) {
		t.Errorf("%+v != %+v", entries, expectedEntries)
	}
}

func TestMultipleGroups(t *testing.T) {
	input := `[foobar]
x = 23
y = 42

[spam]
eggs = moo
ham = oink`
	expectedEntries := Entries{
		"foobar": Group{
			Entry{"x", "23"},
			Entry{"y", "42"}},
		"spam" : Group{
			Entry{"eggs", "moo"},
			Entry{"ham", "oink"}}}
	entries, err := ParseDesktopEntryString(input)
	if err != nil {
		t.Errorf("error: %s", err)
	}
	if !reflect.DeepEqual(entries, expectedEntries) {
		t.Errorf("%+v != %+v", entries, expectedEntries)
	}
}

func TestFile(t *testing.T) {
	entries, err := ParseDesktopEntryFile("example.desktop")
	if err != nil {
		t.Errorf("error: %s", err)
	}
	expectedEntries := Entries{
		"Desktop Entry": Group{
			Entry{"Version", "1.0"},
			Entry{"Type", "Application"},
			Entry{"Name", "Foo Viewer"},
			Entry{"Comment", "The best viewer for Foo objects available!"},
			Entry{"TryExec", "fooview"},
			Entry{"Exec", "fooview %F"},
			Entry{"Icon", "fooview.png"},
			Entry{"MimeType", "image/x-foo;"},
			Entry{"X-KDE-Library", "libfooview"},
			Entry{"X-KDE-FactoryName", "fooviewfactory"},
			Entry{"X-KDE-ServiceType", "FooService"}}}
	if !reflect.DeepEqual(entries, expectedEntries) {
		t.Errorf("%+v != %+v", entries, expectedEntries)
	}
}