package desktopentry

import (
	"fmt"
	"strings"
	"errors"
)

// this is the printable part of the ASCII characters minus the brackets [ and ]
const allowedGroupHeaderChars = " !\"#$%&'()*+,-./0123456789:;<=>?@" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ\\^_`abcdefghijklmnopqrstuvwxyz{|}~"

type Entry struct {
	Key, Value string
}

type Group []Entry
//type Entries []Entry

// key is the group header
//type Group map[string]Entries

//type DesktopEntry []Group
// FIXME: ponder about the following line!
type DesktopEntry map[string]Group

// FIXME: this function is fundamentally broken! FIX IT!
func ParseDesktopEntry(input string) (DesktopEntry, error) {
	// iterate over each line
	// if a group header occurs, make a new group and collect its entries
	// until another group header occurs. repeat util EOF or error.
	dEntry := DesktopEntry{}
	group := Group{}
	groupName := ""
	lines := strings.Split(input, "\n")
	for lineno, line := range lines {
		switch {
		case isGroupHeader(line):
			// add all gathered entries to the yet-current group if
			// this line does not introduce the *first* group
			if groupName != "" {
				dEntry[groupName] = group
			}
			// remove the brackets from the line to get the group name
			groupName = line[1:len(line)-1]
			group = Group{}
		//case strings.HasPrefix(line, "#"), line == "":
		case strings.HasPrefix(line, "#"), strings.TrimSpace(line) == "":
			// line is a comment or an empty line, ignore it
			//break
		default:
			// line is not a group header -> must be key=value
			// form. otherwise it's a parsing error
			entry, err := parseKeyValue(line)
			if err != nil {
				// not a valid line
				errmsg := fmt.Sprintf(
					"invalid input line: %s (line %d)",
					line, lineno)
				return DesktopEntry{}, errors.New(errmsg)
			}
			if groupName == "" {
				// key-value line is valid, but the group
				// header is missing
				errmsg := fmt.Sprintf(
					"missing group header before " +
					"entry declarations (line %d)", lineno)
				return DesktopEntry{}, errors.New(errmsg)
			}
			group = append(group, entry)
			//group[groupName] = append(group[groupName], entry)
		}
	}
	return dEntry, nil
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
