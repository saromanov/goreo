package main

import (
	"log"
	"os"

	"github.com/saromanov/goreo/internal/builder"
	"github.com/urfave/cli"
)

const (
	release = "release"
)

func main() {
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  release,
			Usage: "releasing of the project",
		},
	}

	app.Action = func(c *cli.Context) error {
		if c.Bool(release) {
			builder.Build()
		}
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatalf("unable to init argument parser: %v", err)
	}

}