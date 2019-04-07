package main

import (
	"mycontainer"
	log "github.com/Sirupsen/logrus"
	"os"
)

func Run(tty bool, command string) {
	parent := mycontainer.NewParentProcess(tty, command)
	if err := parent.Start(); err != nil {
		log.Error(err)
	}
	parent.Wait()
	os.Exit(-1)
}
