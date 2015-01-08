package main

import (
	"os"
)

func main() {
	tft := &TFT{outStream: os.Stdout, errStream: os.Stderr}
	os.Exit(tft.Run(os.Args))
}
