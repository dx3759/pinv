package main

import (
	"embed"
	"log"
	"os"

	cli "github.com/urfave/cli/v2"
	"github.com/yzimhao/pinv"
)

//go:embed templates/*
var emfs embed.FS

func main() {

	app := &cli.App{
		Name:  "pinv",
		Usage: "pinv is a tool for managing your files",
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
			pinv.GloOptions.Host = c.String("host")
			pinv.GloOptions.Port = c.Int("port")
			pinv.GloOptions.RootDir = c.String("root")
			pinv.Run(emfs)
			return nil
		},
		Commands: []*cli.Command{
			{
				Name:    "version",
				Aliases: []string{"c"},
				Usage:   "show version",
				Action: func(c *cli.Context) error {
					pinv.ShowVersion()
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
