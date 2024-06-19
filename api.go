package main

import (
	"encoding/json"
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
