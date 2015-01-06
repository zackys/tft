package main

import (
	"github.com/codegangsta/cli"
	"github.com/zackys/go.p/encoding"
	"github.com/zackys/go.p/file"
	"github.com/zackys/go.p/file/text"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "tft"
	app.Version = "0.1"
	app.Usage = ""
	app.Author = "zackys"
	app.Email = "zackys.github@gmail.com"
	app.Commands = Commands
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "j",
			Usage: "出力コード：JIS",
		},
		cli.BoolFlag{
			Name:  "e",
			Usage: "出力コード：EUCJP",
		},
		cli.BoolFlag{
			Name:  "s",
			Usage: "出力コード：SJIS",
		},
		cli.BoolFlag{
			Name:  "w",
			Usage: "出力コード：UTF8（BOMなし）",
		},
		cli.BoolFlag{
			Name:  "w16",
			Usage: "出力コード：UTF16（BOMなし）",
		},
		//		cli.BoolFlag {
		//			Name:  "w8",
		//			Usage: "出力コード：UTF8",
		//		},
	}

	app.Action = func(c *cli.Context) {
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
		println(enc.String())
		println("---")
		txt := text.New(enc)
		txt.ReadFrom(b)
		txt.WriteTo(os.Stdout, encoding.UTF8)
	}

	app.Run(os.Args)
}

func newText(c *cli.Context) *text.Text {
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

	return txt
}

func outEnc(c *cli.Context) encoding.Encoder {
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

	return outEnc
}

func fout(c *cli.Context) *os.File {
	outFname := c.GlobalString("O")

	var fout *os.File
	var err error

	if len(outFname) < 1 {
		fout = os.Stdout
	} else {
		fout, err = os.Create(outFname)
		if err != nil {
			panic(err)
		}
	}

	return fout
}
