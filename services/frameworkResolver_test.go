package services

import (
	"testing"

	. "github.com/franela/goblin"
)

func Test_FrameworkResolver(t *testing.T) {
	g := Goblin(t)
	g.Describe("Framework Resolver", func() {

		g.It("should return 'generic' from DetermineFrameworks since its the first file in test-files", func() {
			frameworks := map[string]string{
				"VERSION":      "generic",
				"package.json": "node",
			}
			directory := "../test-files"

			expected := "generic"

			actual, _ := DetermineFrameworks(directory, frameworks)
			g.Assert(actual).Equal(expected)
		})

		g.It("should return 'node' from DetermineFrameworks if its the only matching framework", func() {
			frameworks := map[string]string{
				"package.json": "node",
			}
			directory := "../test-files"

			expected := "node"

			actual, _ := DetermineFrameworks(directory, frameworks)
			g.Assert(actual).Equal(expected)
		})

		g.It("should return an error if no frameworks are matched", func() {
			frameworks := map[string]string{
				"foobar.txt": "Foobar Framework",
			}
			directory := "../test-files"

			expected := "No frameworks matched"

			_, err := DetermineFrameworks(directory, frameworks)
			actual := err.Error()

			g.Assert(actual).Equal(expected)
		})

	})
}
