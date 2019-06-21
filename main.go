package main

import (
	"github.com/labstack/echo/v4"

	"github.com/kaz/flos/audit"
	"github.com/kaz/flos/beacon"
	"github.com/kaz/flos/messaging"
	"github.com/kaz/flos/power"
	"github.com/kaz/flos/proxy"
	"github.com/kaz/flos/state"
)

func main() {
	e := echo.New()
	e.HideBanner = true
	e.HTTPErrorHandler = messaging.ErrorHandler

	e.Pre(proxy.Middleware)
	e.Pre(messaging.Middleware)

	power.StartService(e.Group("/power"))
	state.StartService(e.Group("/state"))
	beacon.StartService(e.Group("/beacon"))

	go audit.StartWorker()

	e.Logger.Fatal(e.Start(power.LISTEN))
}
