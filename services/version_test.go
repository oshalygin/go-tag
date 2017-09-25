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

		g.It("should read the package.json file and return the version found inside", func() {
			file = "../test-files/package.json"

			expected := "1.3.8"

			actual, _ := node()
			g.Assert(actual).Equal(expected)
		})

		g.It("should throw an error if the package.json version is empty", func() {
			file = "../test-files/empty_version_package.json"

			expected := "No date in the version property of package.json"

			_, err := node()

			actual := err.Error()
			g.Assert(actual).Equal(expected)
		})

		g.It("should throw an error if the package.json file is malformed", func() {
			file = "../test-files/NON_EXISTENT_VERSION_FILE"

			expected := "Error reading the package.json file, malformed"

			_, err := node()

			actual := err.Error()
			g.Assert(actual).Equal(expected)
		})

		g.It("should return the version file if the framework is 'node'", func() {
			file = "../test-files/package.json"
			framework := "node"

			expected := "1.3.8"

			actual, _ := GetVersion(framework)
			g.Assert(actual).Equal(expected)
		})

		g.It("should return the version file if the framework is 'generic'", func() {
			file = "../test-files/VERSION"
			framework := "generic"

			expected := "2.4.9"

			actual, _ := GetVersion(framework)
			g.Assert(actual).Equal(expected)
		})

		g.It("should return an error if the framework version is not supported", func() {
			framework := "python"
			expected := "No matching frameworks"

			_, err := GetVersion(framework)

			actual := err.Error()
			g.Assert(actual).Equal(expected)
		})

	})
}
