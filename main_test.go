package main

import (
	"testing"

	. "github.com/franela/goblin"
)

func Test_Main(t *testing.T) {
	g := Goblin(t)
	g.Describe("Main Application", func() {

		g.It("should run successfully without any errors", func() {
			main()
		})

	})
}
