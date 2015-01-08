package main

import (
	"fmt"
	"github.com/codegangsta/cli"
)

var commandAddln = cli.Command{
	Name:  "addln",
	Usage: "",
	Description: `
`,
	Flags: []cli.Flag{
		cli.IntFlag{
			Name:  "start",
			Value: 100000,
			Usage: "start of LineNumber",
		},
		cli.IntFlag{
			Name:  "step",
			Value: 10,
			Usage: "step of LineNumber",
		},
		cli.StringFlag{
			Name:  "format",
			Value: "%06d%s",
			Usage: "format of LineNumber. %s is the body of line.",
		},
	},
	Action: doAddln,
}

func doAddln(c *cli.Context) {

	txt := newText(c)

	tf := newTfAddln(c)
	txt.Transform(tf)

	outEnc := outEnc(c)

	fout := fout(c)
	defer fout.Close()

	txt.WriteTo(fout, outEnc)
}

type tfAddln struct {
	step   int
	format string

	cnt int
}

func newTfAddln(c *cli.Context) *tfAddln {
	ret := &tfAddln{}
	ret.cnt = c.Int("start")
	ret.step = c.Int("step")
	ret.format = c.String("format")

	return ret
}

func (c *tfAddln) Transform(line string) (string, error) {
	ret := fmt.Sprintf(c.format, c.cnt, line)
	c.cnt++
	return ret, nil
}
