package api

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	log "github.com/sirupsen/logrus"
	miscController "wotracker-back/internal/app/api/controller"
)

func registerRoutes(app *iris.Application) {
	log.Debug("registering misc routing")
	miscApi := app.Party("/misc")
	m := mvc.New(miscApi)
	m.Handle(new(miscController.MiscController))
}
