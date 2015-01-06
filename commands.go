package main

import (
	"bufio"
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/zackys/go.p/encoding"
	"github.com/zackys/go.p/file"
	"github.com/zackys/go.p/file/text"
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
	args := c.Args()
	var err error
	var fin *os.File
	if len(args) == 0 {
		fin = os.Stdin
	} else {
		fin, err = os.Open(args[0])
		if err != nil {
			panic(err)
		}
		defer fin.Close()
	}

	b := file.NewBytes()
	err = b.ReadFrom(fin)
	if err != nil {
		panic(err)
	}
	enc := b.SearchEncoding()

	txt := text.New(enc)
	txt.ReadFrom(b)

	itr := txt.Iterator()
	cnt := c.Int("start")
	step := c.Int("step")
	format := c.String("format")

	var outEnc encoding.Encoding = encoding.UTF8
	switch {
	case c.GlobalBool("s"):
		outEnc = encoding.ShiftJIS
	case c.GlobalBool("e"):
		outEnc = encoding.EUCJP
	case c.GlobalBool("j"):
		outEnc = encoding.ISO2022JP
	case c.GlobalBool("w16"):
		outEnc = encoding.UTF16BE
	}

	outFname := c.GlobalString("O")

	var fout *os.File

	if len(outFname) < 1 {
		fout = os.Stdout
	} else {
		fout, err = os.Create(outFname)
		if err != nil {
			panic(err)
		}
		defer fout.Close()
	}
	writer := bufio.NewWriter(fout)

	for itr.HasNext() {
		line := itr.Next()
		ret := fmt.Sprintf(format, cnt, line)
		cnt += step

		bb, _ := outEnc.Encode(ret)
		writer.Write(bb)
	}
	writer.Flush()

	//txt.WriteTo(os.Stdout)
}
