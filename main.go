package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lxzan/gws"
)

func main() {
	// web server
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))

	// ws server
	upgrader  := gws.NewUpgrader(&WebSocketHandler{}, &gws.ServerOption{
		ParallelEnabled:  true,                                 
		Recovery:          gws.Recovery,                         
		PermessageDeflate: gws.PermessageDeflate{Enabled: true}, 
	})

	// default ws route
	e.GET("/ws", func(c echo.Context) error {
		socket, err := upgrader.Upgrade(c.Response().Writer, c.Request());

		if err != nil {
			return err
		}
		go func() {
			socket.ReadLoop() 
		}()
		return nil
	})
}