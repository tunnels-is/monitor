package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/lxzan/gws"
)

func main() {
	// web server
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})


	// ws server for react
	upgrader  := gws.NewUpgrader(&WebSocketHandler{}, &gws.ServerOption{
		ParallelEnabled:  true,                                 
		Recovery:          gws.Recovery,                         
		PermessageDeflate: gws.PermessageDeflate{Enabled: true}, 
	})

	// ws server for log reciver route
	logUpgrader := gws.NewUpgrader(&LogReceiverSocketHandler{}, &gws.ServerOption{
		ParallelEnabled:  true,                                 
		Recovery:          gws.Recovery,                         
		PermessageDeflate: gws.PermessageDeflate{Enabled: true}, 
	})

	// default ws route
	e.GET("/connect", func(c echo.Context) error {
		socket, err := upgrader.Upgrade(c.Response().Writer, c.Request());

		if err != nil {
			return err
		}
		go func() {
			socket.ReadLoop() 
		}()
		return nil
	})

	// log reciver ws route
	e.GET("/v1/json/dynamic", func(c echo.Context) error {
		socket, err := logUpgrader.Upgrade(c.Response().Writer, c.Request());

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