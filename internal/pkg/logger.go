package logger

import (
	"github.com/kataras/iris/v12"
	log "github.com/sirupsen/logrus"
	"time"
)

func Logger() iris.Handler {
	return func(ctx iris.Context) {
		t := time.Now()
		ctx.Values().Set("begin", t.Format(time.UnixDate))
		// Set a shared variable between handlers
		ctx.Values().Set("framework", "iris")

		// before request
		ctx.Next()

		// after request
		latency := time.Since(t)

		// access the status we are sending
		status := ctx.GetStatusCode()
		log.Debugf("status code: %d \texecution time: %s \t route: %s", status, latency, ctx.Request().URL)

	}
}
