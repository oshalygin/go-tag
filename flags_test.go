package main

import (
	"io/ioutil"
	"testing"

	. "github.com/franela/goblin"
)

func Test_Flags(t *testing.T) {
	g := Goblin(t)
	g.Describe("Flags", func() {

		g.It("should read the local VERSION file and print it", func() {
			content, _ := ioutil.ReadFile("./VERSION")
			expected := string(content)

			actual := printVersion()
			g.Assert(actual).Equal(expected)
		})

	})
}
