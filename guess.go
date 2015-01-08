package main

import (
	"github.com/codegangsta/cli"

	"os"
)

var commandGuess = cli.Command{
	Name:  "guess",
	ShortName: "g",
	Usage: "",
	Description: `
`,
	Action: doGuess,
}

func doGuess(c *cli.Context) {
	txt := newText(c)

	os.Stdout.WriteString(txt.Encoding().String())
}