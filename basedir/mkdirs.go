package basedir

import (
	"os"
	"path"
)

// make sure that the resulting path exists and mkdir missing dirs, if
// necessary
func makeEnvPathFunc(envVariable string) func(string) (string, error) {
	return func(envPath string) (string, error) {
		absPath := path.Join(envVariable, envPath)
		return absPath, os.MkdirAll(absPath, 0700)
	}
}

var MakeDataPath = makeEnvPathFunc(XDG_DATA_HOME)

var MakeConfigPath = makeEnvPathFunc(XDG_CONFIG_HOME)

var MakeCachePath = makeEnvPathFunc(XDG_CACHE_HOME)
