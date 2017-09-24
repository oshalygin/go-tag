package main

import (
	"flag"

	"github.com/fatih/color"
	"github.com/oshalygin/go-tag/constants"
	"github.com/oshalygin/go-tag/services"
)

var versionFlagPassed bool
var directory = "."

func init() {
	flag.BoolVar(&versionFlagPassed, "version", false, "Print the version of the application")
	flag.Parse()
}

func main() {
	if versionFlagPassed == true {
		printVersion()
	}

	framework, err := services.DetermineFrameworks(directory, constants.Frameworks)

	if err != nil {
		color.Red("No framework files matched")
		return
	}

	color.Green(framework)

}
