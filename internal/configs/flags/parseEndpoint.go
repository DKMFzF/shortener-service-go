package flags

import "flag"

var (
	EndpointFlag      = flag.String("endpoint", "", "flag for endpoing")
	EndpointShortFlag = flag.String("e", "", "flag for endpoing short")
)

func ParseEndpoint() string {
	flag.Parse()
	if *EndpointFlag == "" {
		return *EndpointShortFlag
	}
	return *EndpointFlag
}
