package services

import (
	"os"
	"os/exec"
	"testing"

	. "github.com/franela/goblin"
)

// Command spys
var passedInCommand string
var passedInArgs []string

// From Go's source code
// https://golang.org/src/os/exec/exec_test.go
func fakeExecCommand(command string, args ...string) *exec.Cmd {

	passedInCommand = command
	passedInArgs = args

	cs := []string{"-test.run=TestHelperProcess", "--", command}
	cs = append(cs, args...)
	cmd := exec.Command(os.Args[0], cs...)
	cmd.Env = []string{"GO_WANT_HELPER_PROCESS=1"}
	return cmd
}

func Test_GitService(t *testing.T) {
	g := Goblin(t)
	g.Describe("Git Service", func() {

		g.It("should succeed when tagging the current repository with a version", func() {

			execCommand = fakeExecCommand
			defer func() {
				execCommand = exec.Command
			}()

			version := "1.0.3"

			actual := Tag(version)
			g.Assert(actual).Equal(true)
		})

		g.It("should call the proper git command when calling Tag", func() {

			expected := "git"
			execCommand = fakeExecCommand
			defer func() {
				execCommand = exec.Command
			}()

			version := "1.0.3"
			Tag(version)

			actual := passedInCommand
			g.Assert(actual).Equal(expected)
		})

		g.It("should call Tag with the args 'tag' and the version", func() {

			version := "1.0.3"
			expected := []string{"tag", version}

			execCommand = fakeExecCommand
			defer func() {
				execCommand = exec.Command
			}()

			Tag(version)

			actual := passedInArgs
			g.Assert(actual).Equal(expected)
		})

		g.It("should succeed when calling PushTag", func() {

			execCommand = fakeExecCommand
			defer func() { execCommand = exec.Command }()

			version := "1.0.3"
			actual := PushTag(version)

			g.Assert(actual).Equal(true)
		})

		g.It("should call the git command when calling PushTag", func() {

			expected := "git"
			execCommand = fakeExecCommand
			defer func() {
				execCommand = exec.Command
			}()

			version := "1.0.3"
			PushTag(version)

			actual := passedInCommand
			g.Assert(actual).Equal(expected)
		})

		g.It("should call PushTag with the args 'tag' and the version", func() {

			version := "1.0.3"
			expected := []string{"push", "--tags"}

			execCommand = fakeExecCommand
			defer func() {
				execCommand = exec.Command
			}()

			PushTag(version)

			actual := passedInArgs
			g.Assert(actual).Equal(expected)
		})

	})
}
