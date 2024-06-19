package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func UserConfigRecieveHandler(c echo.Context) error {
	var config SocketConfig
	if err := c.Bind(&config); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	key := "exampleKey" 
	SetConnectedSocket(key, ConnectedSocket{Config: config})
	fmt.Println(globalSockets["exampleKey"].Config.Datapoints[0].Axis);
	return c.JSON(http.StatusOK, map[string]string{"status": "success"})
}