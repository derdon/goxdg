package desktopentry

import (
	"fmt"
	"strings"
	"errors"
	"io/ioutil"
)

// this is the printable part of the ASCII characters minus the brackets [ and ]
const allowedGroupHeaderChars = " !\"#$%&'()*+,-./0123456789:;<=>?@" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ\\^_`abcdefghijklmnopqrstuvwxyz{|}~"

type Entry struct {
	Key, Value string
}

type Group []Entry

type Entries map[string]Group

func (entries Entries) String() (out string) {
	for groupName, group := range entries {
		out += fmt.Sprintf("[%s]\n", groupName)
		for _, entry := range group {
			out += fmt.Sprintf("%s = %s\n", entry.Key, entry.Value)
		}
	}
	return strings.TrimSpace(out)
}

// FIXME: this function is fundamentally broken! FIX IT!
func ParseDesktopEntryString(input string) (Entries, error) {
	entries := Entries{}
	group := Group{}
	groupName := ""
	isFirstGroup := true
	lines := strings.Split(input, "\n")
	// iterate over each line
	// if a group header occurs, make a new group and collect its entries
	// until another group header occurs.
	for lineno, line := range lines {
		switch {
		case isGroupHeader(line):
			// add all gathered entries to the yet-current group if
			// this line does not introduce the *first* group
			if groupName != "" {
				entries[groupName] = group
			} else {
				isFirstGroup = false
			}
			// remove the brackets from the line to get the group name
			groupName = line[1:len(line)-1]
			// the spec says that the first group must be named "Desktop Entry"
			if isFirstGroup && groupName != "Desktop Entry" {
				errmsg := fmt.Sprintf(
					"first group name is '%s', must be 'Desktop Entry'.",
					groupName)
				return entries, errors.New(errmsg)
			}
			// make a new empty group where the following found entries will be
			// appended to
			group = Group{}
			// initialize the new group name with this new empty group
			entries[groupName] = group
		//case strings.HasPrefix(line, "#"), line == "":
		case strings.HasPrefix(line, "#"), strings.TrimSpace(line) == "":
			// line is a comment or an empty line, ignore it
			//break
		default:
			// line is not a group header -> must be key=value
			// form. otherwise it's a parsing error
			entry, err := parseKeyValue(line)
			if err != nil {
				return entries, err
			}
			if groupName == "" {
				// key-value line is valid, but the group
				// header is missing
				errmsg := fmt.Sprintf(
					"missing group header before " +
					"entry declarations (line %d)", lineno + 1)
				return entries, errors.New(errmsg)
			}
			group = append(group, entry)
		}
	}
	// if there was only one group, the group entries have to be assigned here
	if groupName != "" {
		entries[groupName] = group
	}
	return entries, nil
}

func ParseDesktopEntryFile(filename string) (entries Entries, err error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	return ParseDesktopEntryString(string(content))
}

func isGroupHeader(line string) bool {
	// A group header is a string which starts with a [ and ends with a ].
	// Between the brackets only the ASCII set is allowed (except the
	// control characters and [ and ]).
	if !strings.HasPrefix(line, "[") || !strings.HasSuffix(line, "]") {
		return false
	}
	groupName := line[1:len(line)-1]
	if groupName == "" {
		return false
	}
	for _, char := range groupName {
		if !strings.ContainsRune(allowedGroupHeaderChars, char) {
			return false
		}
	}
	return true
}

func parseKeyValue(line string) (Entry, error) {
	cleanedLine := strings.TrimSpace(line)
	// XXX: what is the difference between SplitN and SplitAfterN?
	fields := strings.SplitN(cleanedLine, "=", 2)
	if len(fields) < 2 {
		errmsg := "invalid entry line: missing assignment operator ="
		return Entry{}, errors.New(errmsg)
	}
	key := strings.TrimSpace(fields[0])
	// the doc says: "Key names must contain only the characters A-Za-z0-9-"
	if key == "" {
		return Entry{}, errors.New("missing key")
	}
	value := strings.TrimSpace(fields[1])
	if value == "" {
		return Entry{}, errors.New("missing value")
	}
	return Entry{key, value}, nil
}
