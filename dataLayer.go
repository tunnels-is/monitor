package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime/debug"
	"sync"
)

var (
	DiskLock = sync.Mutex{}
	basePath = "./mdata"
	fileType = "json"
)

type FileType string

var (
	DataCenterFile FileType = "datacenter"
	RowFile        FileType = "row"
	RackFile       FileType = "rack"
	UnitFile       FileType = "unit"
	ConnectionFile FileType = "connection"
	NetworkFile    FileType = "network"
)

type StoragePath string

var (
	DataCenterPath StoragePath = "datacenters"
	RowsPath       StoragePath = "rows"
	RackPath       StoragePath = "racks"
	UnitPath       StoragePath = "units"
	ConnectionPath StoragePath = "connections"
	NetworkPath    StoragePath = "networks"
)

func GetObjectPath(p StoragePath, t FileType, id int) string {
	return fmt.Sprintf(
		"%s%s%s%s%s_%d.%s",
		basePath,
		string(os.PathSeparator),
		p,
		string(os.PathSeparator),
		t,
		id,
		fileType,
	)
}

func GetDirPath(p StoragePath, t FileType) string {
	return fmt.Sprintf(
		"%s%s%s%s",
		basePath,
		string(os.PathSeparator),
		p,
		string(os.PathSeparator),
	)
}

func WriteObject(p StoragePath, t FileType, id int, object interface{}) (err error) {
	DiskLock.Lock()
	defer func() {
		r := recover()
		if r != nil {
			log.Println(r, string(debug.Stack()))
		}
		DiskLock.Unlock()
	}()

	dirErr := os.MkdirAll(GetDirPath(p, t), 0o777)
	if dirErr != nil {
		return dirErr
	}

	data, encodingErr := json.Marshal(object)
	if encodingErr != nil {
		return encodingErr
	}

	return os.WriteFile(GetObjectPath(p, t, id), data, 0o777)
}

func GetObject(p StoragePath, t FileType, id int) (data []byte, err error) {
	data, err = os.ReadFile(GetObjectPath(p, t, id))
	if err != nil {
		return nil, err
	}
	return
}

func GetObjectFullPath(path string) (data []byte, err error) {
	data, err = os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return
}

func GetObjects(startPath string, action func(objectPath string)) (err error) {
	walkFunc := func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			action(path)
		}
		return nil
	}

	filepath.WalkDir(startPath, walkFunc)

	return err
}
