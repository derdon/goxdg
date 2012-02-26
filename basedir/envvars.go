package basedir

import "os"

var home = os.Getenv("$HOME")
var local = home + "/.local"
var (
	XDG_DATA_HOME   = getenv("XDG_DATA_HOME", local+"/share")
	XDG_CONFIG_HOME = getenv("XDG_CONFIG_HOME", home+"/.config")
	XDG_DATA_DIRS   = getenv(
		"XDG_DATA_DIRS", "/usr/local/share:/usr/share/")
	XDG_CONFIG_DIRS = getenv("XDG_CONFIG_DIRS", "/etc/xdg")
	XDG_CACHE_HOME  = getenv("XDG_CACHE_HOME", home+"/.cache")
)

func getenv(key, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) > 0 {
		return value
	}
	return defaultValue
}
