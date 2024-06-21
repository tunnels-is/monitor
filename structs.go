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

// DATACENTER STUFF
// DATACENTER STUFF
// DATACENTER STUFF
// DATACENTER STUFF
// DATACENTER STUFF
// DATACENTER STUFF

type Datacenter struct {
	ID   int    `json:"ID"`
	Tag  string `json:"Tag"`
	Info string `json:"Info"`
}

type Row struct {
	ID   int    `json:"ID"`
	Tag  string `json:"Tag"`
	Info string `json:"Info"`

	DCID int `json:"DatacenterID"`
}

type Rack struct {
	ID   int    `json:"ID"`
	Tag  string `json:"Tag"`
	Info string `json:"Info"`

	RowID int `json:"RowID"`

	TotalU       int       `json:"TotalU"`
	BreakerLimit int       `json:"BreakerLimit"`
	PowerType    PowerType `json:"PowerType"`
	// PowerWarning int       `json:"BreakerWarning"`
	// ???

	// Visual
	Color string `json:"Color"`
}

type Unit struct {
	ID   int    `json:"ID"`
	Info string `json:"Info"`
	Tag  string `json:"Tag"`

	RackID int `json:"RackID"`

	BMC      string     `json:"BMC"`
	SSH      string     `json:"SSH"`
	IPs      []string   `json:"IPs"`
	Networks []*Network `json:"Networks"`

	// Visual
	Size  int    `json:"Size"`
	Color string `json:"Color"`
}

type Connection struct {
	ID   int    `json:"ID"`
	Info string `json:"Info"`
	Tag  string `json:"Tag"`

	From int `json:"From"`
	To   int `json:"To"`

	// Visual
	Color string `json:"Color"`
}

type Network struct {
	ID   int    `json:"ID"`
	Info string `json:"Info"`
	Tag  string `json:"Tag"`

	CIDR string `json:"CIDR"`
}

type PowerType string

const (
	AC PowerType = "ac"
	DC PowerType = "dc"
)

// DATACENTER STUFF
// DATACENTER STUFF
// DATACENTER STUFF
// DATACENTER STUFF
// DATACENTER STUFF
// DATACENTER STUFF

// DATACENTER STUFF
// DATACENTER STUFF
// DATACENTER STUFF
// DATACENTER STUFF
// DATACENTER STUFF
// DATACENTER STUFF


// DATACENTER STUFF
// DATACENTER STUFF
// DATACENTER STUFF
// DATACENTER STUFF
// DATACENTER STUFF
// DATACENTER STUFF

type ConnectedSocket struct {
	Config     SocketConfig
}

var (
	globalSocketsM = make(map[string]*ConnectedSocket)
	mutex         sync.Mutex
)

func SetConnectedSocket(key string, socket *ConnectedSocket) {
	mutex.Lock()
	defer mutex.Unlock()
	globalSocketsM[key] = socket
}

func GetConnectedSocket(key string) (*ConnectedSocket, bool) {
	mutex.Lock()
	defer mutex.Unlock()
	socket, exists := globalSocketsM[key]
	return socket, exists
}