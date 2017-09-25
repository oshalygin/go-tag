package services

import (
	"encoding/json"
	"errors"
	"io/ioutil"

	"github.com/oshalygin/go-tag/models"
)

var file string

// GetVersion will read the appropriate file based on the framework and return the version value
func GetVersion(framework string) (string, error) {
	switch framework {

	case "node":
		version, err := node()
		if err != nil {
			return "", err
		}

		return version, nil

	case "generic":

		version, err := generic()
		if err != nil {
			return "", err
		}

		return version, nil

	default:
		return "", errors.New("No matching frameworks")
	}
}

func node() (string, error) {
	if file == "" {
		file = "./package.json"
	}

	result, err := ioutil.ReadFile(file)
	if err != nil {
		return "", errors.New("Error reading the package.json file, malformed")
	}

	var configurationFile models.PackageJSON
	json.Unmarshal(result, &configurationFile)

	version := configurationFile.Version

	if version == "" {
		return "", errors.New("No date in the version property of package.json")
	}

	return version, nil

}

func generic() (string, error) {

	if file == "" {
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
