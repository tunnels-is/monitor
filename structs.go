package main

import "sync"

type SocketConfig struct {
	Datapoints []*Vector
}

type Vector struct {
	Axis      VectorAxis
	Tag       string
	Namespace string
	Type      VectorType
	Interval  VectorInterval
	DataType  VectorDataType
}

type VectorAxis string

const (
	X VectorAxis = "x"
	Y VectorAxis = "y"
	Z VectorAxis = "z"
)

type VectorDataType string

const (
	I VectorDataType = "int"
	F VectorDataType = "float"
)

type VectorType string

const (
	MIN  VectorType = "min"
	MAX  VectorType = "max"
	SUM  VectorType = "sum"
	MEAN VectorType = "mean"
)

type VectorInterval string

const (
	S5  VectorInterval = "5sec"
	S10 VectorInterval = "10sec"
	S30 VectorInterval = "30sec"
	M1  VectorInterval = "1min"
	M5  VectorInterval = "5min"
	M10 VectorInterval = "10min"
	M30 VectorInterval = "30min"
	HR1 VectorInterval = "1hr"
)

type ConnectedSocket struct {
	Config     SocketConfig
}

var (
	globalSockets = make(map[string]ConnectedSocket)
	mutex         sync.Mutex
)

func SetConnectedSocket(key string, socket ConnectedSocket) {
	mutex.Lock()
	defer mutex.Unlock()
	globalSockets[key] = socket
}

func GetConnectedSocket(key string) (ConnectedSocket, bool) {
	mutex.Lock()
	defer mutex.Unlock()
	socket, exists := globalSockets[key]
	return socket, exists
}