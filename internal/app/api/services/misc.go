package misc

import (
	"github.com/kataras/iris/v12"
	"strconv"
	"time"
	health "wotracker-back/internal/app/api/models"
)

func GetHealthService(ctx iris.Context) health.Health {
	begin := ctx.Values().GetString("begin")
	beginFullTime, _ := time.Parse(time.UnixDate, begin)
	elapsedTime := time.Since(beginFullTime)

	myHealth := health.Health{
		Code:         200,
		ResponseTime: strconv.FormatInt(elapsedTime.Milliseconds(), 10) + "ms",
	}
	return myHealth
}
