package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/lxzan/gws"
)

func main() {
	// web serverj
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	wsV1 := e.Group("/v1/ws")
	apiV1 := e.Group("/v1/api")

	apiV1.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	apiV1.GET("/datacenters", ListDatacenters)
	apiV1.GET("/datacenters/:id", FindDatacenter)

	// ws upgrader for react
	upgrader := gws.NewUpgrader(&WebSocketHandler{}, &gws.ServerOption{
		ParallelEnabled:   true,
		Recovery:          gws.Recovery,
		PermessageDeflate: gws.PermessageDeflate{Enabled: true},
	})

	// ws upgrader for log reciver route
	logUpgrader := gws.NewUpgrader(&LogReceiverSocketHandler{}, &gws.ServerOption{
		ParallelEnabled:   true,
		Recovery:          gws.Recovery,
		PermessageDeflate: gws.PermessageDeflate{Enabled: true},
	})

	// default ws route
	wsV1.GET("/connect", func(c echo.Context) error {
		socket, err := upgrader.Upgrade(c.Response().Writer, c.Request())
		if err != nil {
			return err
		}
		go func() {
			socket.ReadLoop()
		}()
		return nil
	})

	// log reciver ws route
	wsV1.GET("json/dynamic", func(c echo.Context) error {
		socket, err := logUpgrader.Upgrade(c.Response().Writer, c.Request())
		if err != nil {
			return err
		}
		go func() {
			socket.ReadLoop()
		}()
		return nil
	})

	e.Logger.Fatal(e.Start(":1323"))
}
