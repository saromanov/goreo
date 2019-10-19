package main

import (
	"fmt"
	"log"
	"os"

	"github.com/saromanov/goreo/internal/config"
	"github.com/saromanov/goreo/internal/pipeline"
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
		conf, err := config.Unmarshal("config.yaml")
		if err != nil {
			return fmt.Errorf("unable to unmarshal config: %v", err)
		}
		if c.Bool(release) {
			pipe := pipeline.New(conf)
			pipe.Run()
		}
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatalf("unable to init argument parser: %v", err)
	}

}
