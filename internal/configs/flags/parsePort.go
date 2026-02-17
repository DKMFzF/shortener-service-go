package flags

import (
	"flag"
)

var (
	CustomAddrFlag      = flag.String("port", "", "custom address and port to run server")
	CustomAddrShortFlag = flag.String("p", "", "custom address and port to run server short")
)

func ParsePort() string {
	flag.Parse()
	if *CustomAddrFlag == "" {
		return *CustomAddrShortFlag
	}
	return *CustomAddrFlag
}
