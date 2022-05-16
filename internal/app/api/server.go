package api

import (
	"github.com/iris-contrib/swagger/v12"
	"github.com/iris-contrib/swagger/v12/swaggerFiles"
	"github.com/kataras/iris/v12"
	log "github.com/sirupsen/logrus"
	_ "wotracker-back/docs"
	logger "wotracker-back/internal/pkg"
)

func StartServer(serverPort string) (app *iris.Application) {
	app = iris.New()
	log.Debug("loading middleware...")
	app.Use(logger.Logger())

	log.Debug("adding swagger")
	config := &swagger.Config{
		URL: "http://localhost:" + serverPort + "/swagger/doc.json", //The url pointing to API definition
	}
	// use swagger middleware to
	app.Get("/swagger/{any:path}", swagger.CustomWrapHandler(config, swaggerFiles.Handler))

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
