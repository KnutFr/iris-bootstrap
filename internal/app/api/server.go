package api

import (
	"github.com/kataras/iris/v12"
	log "github.com/sirupsen/logrus"
	logger "wotracker-back/internal/pkg"
)

func StartServer(serverPort string) (app *iris.Application) {
	app = iris.New()
	log.Debug("loading middleware...")
	app.Use(logger.Logger())
	log.Debug("registering routes")
	registerRoutes(app)

	err := app.Listen(":" + serverPort)
	if err != nil {
		log.Fatalf("cannot start server : %s", err)
	} else {
		log.Infof("starting server on port %s", serverPort)
	}
	return
}
