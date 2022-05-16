package miscController

import (
	"github.com/kataras/iris/v12"
	log "github.com/sirupsen/logrus"
	misc "wotracker-back/internal/app/api/services"
)

type MiscController struct {
}

func (c *MiscController) GetHealth(ctx iris.Context) {
	myHealth := misc.GetHealthService(ctx)
	_, err := ctx.JSON(myHealth)
	if err != nil {
		log.Errorf("cannot encode health: %s", err)
	}
}
