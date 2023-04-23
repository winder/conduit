package data

import (
	"fmt"
	"os"
	"path/filepath"
)

// isDir returns true if the specified directory is valid. Copied from Indexer util.IsDir
func isDir(path string) bool {
	fi, err := os.Stat(path)
	return err == nil && fi.IsDir()
}

// fileExists checks to see if the specified file (or directory) exists. Copied from Indexer util.FileExists
func fileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	fileExists := err == nil
	return fileExists
}

// getConfigFromDataDir Given the data directory, configuration filename and a list of types, see if
// a configuration file that matches was located there.  If no configuration file was there then an
// empty string is returned.  If more than one filetype was matched, an error is returned.
// Copied from Indexer util.GetConfigFromDataDir
func getConfigFromDataDir(dataDirectory string, configFilename string, configFileTypes []string) (string, error) {
	count := 0
	fullPath := ""
	var err error

	for _, configFileType := range configFileTypes {
		autoloadParamConfigPath := filepath.Join(dataDirectory, configFilename+"."+configFileType)
		if fileExists(autoloadParamConfigPath) {
			count++
			fullPath = autoloadParamConfigPath
		}
	}

	if count > 1 {
		return "", fmt.Errorf("config filename (%s) in data directory (%s) matched more than one filetype: %v",
			configFilename, dataDirectory, configFileTypes)
	}

	// if count == 0 then the fullpath will be set to "" and error will be nil
	// if count == 1 then it fullpath will be correct
	return fullPath, err
}
