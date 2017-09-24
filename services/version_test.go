package services

import (
	"testing"

	. "github.com/franela/goblin"
)

func Test_VersionService(t *testing.T) {
	g := Goblin(t)
	g.Describe("Version Service", func() {

		g.It("should read the VERSION file and return the version found inside", func() {
			file = "../test-files/VERSION"

			expected := "2.4.9"

			actual, _ := generic()
			g.Assert(actual).Equal(expected)
		})

		g.It("should throw an error if the VERSION file is empty", func() {
			file = "../test-files/EMPTY_VERSION"

			expected := "No data in VERSION file"

			_, err := generic()

			actual := err.Error()
			g.Assert(actual).Equal(expected)
		})

		g.It("should throw an error if the VERSION file is malformed", func() {
			file = "../test-files/NON_EXISTENT_VERSION_FILE"

			expected := "Error reading VERSION file, malformed"

			_, err := generic()

			actual := err.Error()
			g.Assert(actual).Equal(expected)
		})

	})
}
