package main

import (
	"fmt"

	"./container"
	log "github.com/Sirupsen/logrus"
	"github.com/urfave/cli"
)

var runCommand = cli.Command{
	Name:  "run",
	Usage: `Create a container with namespace and cgroups limit sdocker run -it [command]`,
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "it",
			Usage: "enable tty",
		},
	},
	Action: func(context *cli.Context) error {
		if len(context.Args()) < 1 {
			return fmt.Errorf("Missing container command")
		}
		cmd := context.Args().Get(0)
		tty := context.Bool("it")
		Run(tty, cmd)
		return nil
	},
}
var initCommand = cli.Command{
	Name:  "init",
	Usage: "Init container process run user's process in container. Do not call it outside",
	Action: func(context *cli.Context) error {
		log.Infof("initializing")
		cmd := context.Args().Get(0)
		err := container.RunContainerInitProcess(cmd, nil)
		return err
	},
}
