package main

import (
	"strings"
	"testing"
)

type Test struct {
	Data string
}

func Test_EnsureGetAndWriteFunctionality(t *testing.T) {
	basePath = "./mdata"
	dummy := new(Test)

	table := map[StoragePath]FileType{
		DataCenterPath: DataCenterFile,
		RowsPath:       RowFile,
		RackPath:       RackFile,
		UnitPath:       UnitFile,
		ConnectionPath: ConnectionFile,
		NetworkPath:    NetworkFile,
	}

	for i, v := range table {
		dummy.Data = string(GetObjectPath(i, v, 0))
		err := WriteObject(i, v, 0, dummy)
		if err != nil {
			t.Fatal(err)
		}
	}

	for i, v := range table {
		data, err := GetObject(i, v, 0)
		if err != nil {
			t.Fatal(err)
		}
		x := GetObjectPath(i, v, 0)
		if !strings.Contains(string(data), x) {
			t.Fatalf("does not have path")
		}
	}

	testObject := func(path string) {
		data, err := GetObjectFullPath(path)
		if err != nil {
			t.Fatal(err)
		}

		if !strings.Contains(string(data), path) {
			t.Fatalf("does not have path")
		}
	}

	err := GetObjects(basePath, testObject)
	if err != nil {
		t.Fatal(err)
	}

	// os.RemoveAll(basePath)
}
