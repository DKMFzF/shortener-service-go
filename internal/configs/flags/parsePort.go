package flags

import (
	"flag"
)

var (
	CustomAddr = flag.String("port", "", "custom address and port to run server")
)

func ParsePort() string {
	flag.Parse()
	return *CustomAddr
}
