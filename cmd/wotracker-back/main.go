package main

import (
	log "github.com/sirupsen/logrus"
	"wotracker-back/internal/app/starter"
)

func main() {
	log.SetLevel(log.DebugLevel)
	log.Info("starting API... ")
	starter.RunApi()
	log.Info("ending API... ")
}
