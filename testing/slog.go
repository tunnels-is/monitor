package main

import (
	"log/slog"
	"os"
)

var logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))

type Attributes map[string]interface{}

var (
	NodeTag  = "london-01-01"
	RouterIP = "55.55.55.55"
)

func injectGlobalIntoAttributes(a Attributes) {
	a["RouterIP"] = RouterIP
	a["Tag"] = NodeTag
}

func LogConnectivityError(a Attributes) {
	injectGlobalIntoAttributes(a)
	logger.Error("ConnectivityError", "data", a)
}

func main() {
	LogConnectivityError(Attributes{
		"Local":   "12.123.123.63:3542",
		"Remote":  "1.1.1.1:53",
		"Network": "tcp",
		"Error":   "big error",
		"Port":    20,
	})
}
