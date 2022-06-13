package main

import (
	"embed"
	"log"
	"os"

	cli "github.com/urfave/cli/v2"
	"github.com/yzimhao/ymfile"
)

//go:embed templates/*
var emfs embed.FS

func main() {

	app := &cli.App{
		Name:  "ymfile",
		Usage: "ymfile is a tool for managing your files",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "host",
				Value: "0.0.0.0",
				Usage: "host to listen",
			},
			&cli.Int64Flag{
				Name:  "port",
				Value: 8080,
				Usage: "port to listen",
			},
			&cli.StringFlag{
				Name:  "root",
				Value: "./",
				Usage: "root directory",
			},
		},
		Action: func(c *cli.Context) error {
			ymfile.GloOptions.Host = c.String("host")
			ymfile.GloOptions.Port = c.Int("port")
			ymfile.GloOptions.RootDir = c.String("root")
			ymfile.Run(emfs)
			return nil
		},
		Commands: []*cli.Command{
			{
				Name:    "version",
				Aliases: []string{"c"},
				Usage:   "show version",
				Action: func(c *cli.Context) error {
					ymfile.ShowVersion()
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
