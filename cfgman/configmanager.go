package cfgman

import (
	"os"
	"path"
)

/*
File a particularly named file. A set of search path is
given through parameter paths. If a path in paths is absolute,
this path will be used directly, otherwise, current working path
will be used as root of the relative path to search. Each path
in paths will be used along with current working directory to search
for the filename file.

Returns the first encountered instance of the file, with true indicating
successful found, false indicating a search failure.
*/
func findFile(paths []string, filename string) (string, bool) {

	pwd, _ := os.Getwd()

	searchPathes := append(paths, pwd)
	for _, searchPath := range searchPathes {

		var fileToSearch string
		if path.IsAbs(searchPath) {
			fileToSearch = path.Join(searchPath, filename)
		} else {
			fileToSearch = path.Join(pwd, filename)
		}

		if _, err := os.Stat(fileToSearch); err == nil {
			return fileToSearch, true
		}
	}
	return "", false
}

func findConfigFile(filename string) (string, bool) {
	return findFile([]string{"cfg", "config", "settings"}, filename)
}

type EurekaSettingsConfiguration struct {
	EurekaServiceSettingsFile   string
	ServiceInstanceSettingsFile string
}

var EurekaSettingsFiles = EurekaSettingsConfiguration{"eureka-service.json", "service-instance.json"}
