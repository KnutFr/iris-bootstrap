package misc

import (
	"github.com/kataras/iris/v12"
	health "wotracker-back/internal/app/api/models"
)

func GetHealthService(ctx iris.Context) health.Health {
	myHealth := health.Health{
		Code:         200,
		ResponseTime: ctx.Values().GetString("latency"),
	}
	return myHealth
}
