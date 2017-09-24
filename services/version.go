package services

import (
	"errors"
	"io/ioutil"
)

var file string

// GetVersion will read the appropriate file based on the framework and return the version value
func GetVersion(framework string) (string, error) {
	switch framework {
	case "node":
		return node(), nil
	case "generic":

		version, err := generic()
		if err != nil {
			return version, nil
		}

		return "", err
	default:
		return "", errors.New("No matching frameworks")
	}
}

func node() string {
	return ""
}

func generic() (string, error) {

	if file == "file" {
		file = "./VERSION"
	}

	result, err := ioutil.ReadFile(file)
	if err != nil {
		return "", errors.New("Error reading VERSION file, malformed")
	}

	content := string(result)
	if len(content) <= 0 {
		return "", errors.New("No data in VERSION file")
	}

	return content, nil
}
