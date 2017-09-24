package main

import (
	"flag"
)

var versionFlagPassed bool

func init() {
	flag.BoolVar(&versionFlagPassed, "version", false, "Print the version of the application")
	flag.Parse()
}

func main() {
	if versionFlagPassed == true {
		printVersion()
	}
}
