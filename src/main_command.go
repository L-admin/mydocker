package main

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/urfave/cli"
	"mycontainer"
)

var runCommand = cli.Command{
	Name: "run",
	Usage: `Create a mycontainer with namespace and cgroups limit
			mydocker run -ti [command]`,
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "ti",
			Usage: "enable tty",
		},
	},
	Action: func(context *cli.Context) error {
		if len(context.Args()) < 1 {
			return fmt.Errorf("Missing mycontainer command")
		}
		cmd := context.Args().Get(0)
		log.Infof("runCommand command: %s", cmd)
		tty := context.Bool("ti")
		Run(tty, cmd)
		return nil
	},
}

var initCommand = cli.Command{
	Name:  "init",
	Usage: "Init mycontainer process run user's process in mycontainer. Do not call it outside",
	Action: func(context *cli.Context) error {
		log.Infof("init come on")
		cmd := context.Args().Get(0)
		log.Infof("initCommand command: %s", cmd)
		err := mycontainer.RunContainerInitProcess(cmd, nil)
		return err
	},
}
