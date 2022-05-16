package logger

import (
	"github.com/kataras/iris/v12"
	log "github.com/sirupsen/logrus"
	"time"
)

func Logger() iris.Handler {
	return func(ctx iris.Context) {
		t := time.Now()
		ctx.Values().Set("begin", t)
		// Set a shared variable between handlers
		ctx.Values().Set("framework", "iris")

		// before request
		ctx.Next()

		// after request
		latency := time.Since(t)

		// access the status we are sending
		status := ctx.GetStatusCode()
		log.Printf("status code: %d, execution time: %s", status, latency)

	}
}
