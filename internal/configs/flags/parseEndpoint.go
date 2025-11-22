package flags

import "flag"

var (
	EndpointFlag = flag.String("endpoint", "", "flag for endpoing")
)

func ParseEndpoint() string {
	flag.Parse()
	return *EndpointFlag
}
