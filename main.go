package main

import (
	"os"

	"github.com/saromanov/goreo/internal/config"
	"github.com/saromanov/goreo/internal/pipeline"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "release",
			Usage: "releasing of the project",
		},
	}

	app.Action = func(c *cli.Context) error {
		conf, err := config.Unmarshal("goreo.yml")
		if err != nil {
			log.Fatalf("unable to unmarshal config: %v", err)
		}
		pipe := pipeline.New(conf)
		if err := pipe.Run(); err != nil {
			log.Fatalf("unable to finish pipeline: %v", err)
		}
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatalf("unable to init argument parser: %v", err)
	}

}
