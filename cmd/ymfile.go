package main

import (
	"embed"
	"flag"

	"github.com/yzimhao/ymfile"
)

//go:embed templates/*
var emfs embed.FS

func main() {
	version := flag.Bool("version", false, "show version")
	flag.Parse()

	if *version {
		ymfile.ShowVersion()
	}
	ymfile.Run(emfs)
}
