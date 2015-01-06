package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"log"
	"os"
)

var Commands = []cli.Command{
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
			Usage: "format of LineNumber",
		},
	},
	Action: doAddln,
}

func debug(v ...interface{}) {
	if os.Getenv("DEBUG") != "" {
		log.Println(v...)
	}
}

func assert(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func doRmln(c *cli.Context) {
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
