package basedir

import (
	"strings"
	"os"
)

var home = os.Getenv("$HOME")
var local = home + "/.local"
var (
	XDG_DATA_HOME   = getOSEnv("XDG_DATA_HOME", local+"/share")
	XDG_CONFIG_HOME = getOSEnv("XDG_CONFIG_HOME", home+"/.config")
	XDG_DATA_DIRS   = getOSEnv(
		"XDG_DATA_DIRS", "/usr/local/share:/usr/share/")
	XDG_CONFIG_DIRS = getOSEnv("XDG_CONFIG_DIRS", "/etc/xdg")
	XDG_CACHE_HOME  = getOSEnv("XDG_CACHE_HOME", home+"/.cache")
)

func makeMap(keyValueForms []string) environ {
	m := make(environ)
	for _, item := range keyValueForms {
		splittedItem := strings.Split(item, "=")
		key, value := splittedItem[0], splittedItem[1]
		m[key] = value
	}
	return m
}

func getOSEnv(key, defaultValue string) string {
	return getEnv(makeMap(os.Environ()), key, defaultValue)
}

func getEnv(mapping environ, key, defaultValue string) string {
	value, exists := mapping[key]
	if exists {
		return value
	}
	return defaultValue
}
