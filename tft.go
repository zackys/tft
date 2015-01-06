package main

import (
	"os"
	"github.com/zackys/go.p/file"
	"github.com/zackys/go.p/file/text"
	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "tft"
	app.Version = "0.1"
	app.Usage = ""
	app.Author = "zackys"
	app.Email = "zackys.github@gmail.com"
	app.Commands = Commands
	app.Flags = []cli.Flag {
		cli.BoolFlag {
			Name:  "j",
			Usage: "出力コード：JIS",
		},
		cli.BoolFlag {
			Name:  "e",
			Usage: "出力コード：EUCJP",
		},
		cli.BoolFlag {
			Name:  "s",
			Usage: "出力コード：SJIS",
		},
		cli.BoolFlag {
			Name:  "w",
			Usage: "出力コード：UTF8（BOMなし）",
		},
		cli.BoolFlag {
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
		txt.WriteTo(os.Stdout)
	}

	app.Run(os.Args)
}
