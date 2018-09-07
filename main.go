package main

import (
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/urfave/cli"
)

const usage = `sdocker is a simple container runtime implementation.`

func main() {
	app := cli.NewApp()
	app.Name = "sdocker"
	app.Usage = usage

	app.Commands = []cli.Command{
		initCommand,
		runCommand,
	}

	app.Before = func(context *cli.Context) error {
		log.SetFormatter(&log.TextFormatter{
			FullTimestamp: true,
		})
		log.SetOutput(os.Stdout)
		return nil
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
