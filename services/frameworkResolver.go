package services

import (
	"errors"
	"io/ioutil"
)

// DetermineFrameworks determines which files are in the root and returns the frameworks
func DetermineFrameworks(directory string, frameworks map[string]string) (string, error) {

	files, _ := ioutil.ReadDir(directory)
	for _, file := range files {

		fileName := file.Name()
		framework, exists := frameworks[fileName]

		if exists {
			return framework, nil
		}
	}

	return "", errors.New("No frameworks matched")
}
