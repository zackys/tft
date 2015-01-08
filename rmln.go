package main

import (
	"github.com/codegangsta/cli"
)

var Commands = []cli.Command{
	commandGuess,
	commandRmln,
	commandAddln,
}

var commandRmln = cli.Command{
	Name:  "rmln",
	Usage: "",
	Description: `
`,
	Action: doRmln,
}

func doRmln(c *cli.Context) {
}


