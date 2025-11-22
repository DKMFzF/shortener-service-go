package flags

import (
	"flag"
	"fmt"
	"os"
)

var (
	EnvFlags    = flag.NewFlagSet("env", flag.ExitOnError)
	DefaultAddr = EnvFlags.String("default", ":8080", "default address and port to run server")
	CustomAddr  = EnvFlags.String("custom", "", "custom address and port to run server")
)

func ParseFlags() string {
	if len(os.Args) < 2 {
		fmt.Println("subcommand required: cnv")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "env":
		EnvFlags.Parse(os.Args[2:])
	default:
		fmt.Println("unknown subcommand:", os.Args[1])
		os.Exit(1)
	}

	if *CustomAddr != "" {
		return *CustomAddr
	}

	return *DefaultAddr
}
