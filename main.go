package main

import (
	"os"
	"log"
)

func main() {
	tft := &TFT{outStream: os.Stdout, errStream: os.Stderr}
	os.Exit(tft.Run(os.Args))
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
