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

func ListRows(c echo.Context) error {
	rows := make([]*Row, 0)

	var errs []error
	action := func(path string) {
		if !strings.Contains(path, "row_") {
			return
		}
		data, err := GetObjectFullPath(path)
		dc := new(Row)
		err = json.Unmarshal(data, dc)
		if err != nil {
			errs = append(errs, err)
			return
		}
		rows = append(rows, dc)
	}

	err := GetObjects(basePath, action)
	if err != nil {
		return c.JSON(500, err)
	}

	if len(errs) != 0 {
		return c.JSON(500, errs)
	}

	return c.JSON(200, rows)
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

func WriteErrorResponse(c echo.Context, code int, err error, msg string) error {
	return c.JSON(code, map[string]string{"msg": msg, "error": err.Error()})
}

func UserConfigRecieveHandler(c echo.Context) error {
	var config SocketConfig
	if err := c.Bind(&config); err != nil {
		return WriteErrorResponse(c, http.StatusBadRequest, err, "Invalid input")
	}

	key := "exampleKey"
	SetConnectedSocket(key, &ConnectedSocket{Config: config})
	fmt.Println(globalSocketsM["exampleKey"].Config.Datapoints[0].Axis)
	return c.JSON(http.StatusOK, map[string]string{"status": "success"})
}
