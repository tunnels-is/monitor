package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

func ListDatacenters(c echo.Context) error {
	dataCenters := make([]*Datacenter, 0)

	var errs []error
	action := func(path string) {
		if !strings.Contains(path, "datacenter_") {
			return
		}
		data, err := GetObjectFullPath(path)
		dc := new(Datacenter)
		err = json.Unmarshal(data, dc)
		if err != nil {
			errs = append(errs, err)
			return
		}
		dataCenters = append(dataCenters, dc)
	}

	err := GetObjects(basePath, action)
	if err != nil {
		return c.JSON(500, err)
	}

	if len(errs) != 0 {
		return c.JSON(500, errs)
	}

	return c.JSON(200, dataCenters)
}

func FindDatacenter(c echo.Context) error {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(400, err)
	}

	data, err := GetObject(DataCenterPath, DataCenterFile, idInt)
	if err != nil {
		return c.JSON(500, err)
	}

	DC := new(Datacenter)
	err = json.Unmarshal(data, DC)
	if err != nil {
		return c.JSON(500, err)
	}

	return c.JSON(200, DC)
}



func ListDatacenters(c echo.Context) error {
	dataCenters := make([]*Datacenter, 0)

	var errs []error
	action := func(path string) {
		if !strings.Contains(path, "datacenter_") {
			return
		}
		data, err := GetObjectFullPath(path)
		dc := new(Datacenter)
		err = json.Unmarshal(data, dc)
		if err != nil {
			errs = append(errs, err)
			return
		}
		dataCenters = append(dataCenters, dc)
	}

	err := GetObjects(basePath, action)
	if err != nil {
		return c.JSON(500, err)
	}

	if len(errs) != 0 {
		return c.JSON(500, errs)
	}

	return c.JSON(200, dataCenters)
}

func FindDatacenter(c echo.Context) error {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(400, err)
	}

	data, err := GetObject(DataCenterPath, DataCenterFile, idInt)
	if err != nil {
		return c.JSON(500, err)
	}

	DC := new(Datacenter)
	err = json.Unmarshal(data, DC)
	if err != nil {
		return c.JSON(500, err)
	}

	return c.JSON(200, DC)
}


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