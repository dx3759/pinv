package main

import (
	"flag"

	"github.com/yzimhao/ymfile"
)

func main() {
	version := flag.Bool("version", false, "show version")
	flag.Parse()

	if *version {
		ymfile.ShowVersion()
	}
	ymfile.Run()
}
