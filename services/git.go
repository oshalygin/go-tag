package services

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"

	"github.com/fatih/color"
)

var execCommand = exec.Command
var directory = "."

func printOutput(buffer ...*bytes.Buffer) {

	for _, output := range buffer {
		if output != nil {
			fmt.Print(string(output.Bytes()))
		}
	}
}

// TagAndPush unifies both Tag and PushTag in a single call
func TagAndPush(version string) error {
	successfullyTagged := Tag(version)

	if !successfullyTagged {
		return errors.New("Failed to Tag the repository")
	}

	successfullyPushedTag := PushTag(version)

	if !successfullyPushedTag {
		return errors.New("Failed to push the Tag to your remote origin")
	}

	return nil
}

// Tag will tag the repository with the passed in version
func Tag(version string) bool {

	color.Blue("[APPLY]: Tag: %s", version)

	commandName := "git"
	args := []string{"tag", version}

	command := execCommand(commandName, args...)
	command.Dir = directory

	commandOutput := &bytes.Buffer{}
	commandErrorOutput := &bytes.Buffer{}
	command.Stdout = commandOutput
	command.Stderr = commandErrorOutput

	err := command.Run()

	printOutput(commandOutput, commandErrorOutput)

	if err != nil {
		fmt.Println(err.Error())
		color.Red("[ERROR]: Could not tag the repository")
		return false
	}

	color.Green("[SUCCESS]: Tagged: %s", version)
	return true

}

// PushTag will push the tags to the remote
func PushTag(version string) bool {

	color.Blue("[PUSH] Tags: %s", version)

	commandName := "git"
	args := []string{"push", "--tags"}

	command := execCommand(commandName, args...)
	command.Dir = directory

	commandOutput := &bytes.Buffer{}
	commandErrorOutput := &bytes.Buffer{}
	command.Stdout = commandOutput
	command.Stderr = commandErrorOutput

	err := command.Run()

	printOutput(commandOutput, commandErrorOutput)

	if err != nil {
		fmt.Println(err.Error())
		color.Red("[ERROR]: Could not push the tags to the remote")
		return false
	}

	color.Green("[SUCCESS]: Pushed Tags: %s", version)
	return true

}
