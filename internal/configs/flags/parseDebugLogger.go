package flags

import "flag"

var (
	IsDebugLoggerFlag      = flag.Bool("debugLogger", false, "flag for debug level logger")
	IsDebugLoggerShortFlag = flag.Bool("dl", false, "Shorhand for --debugLogger")
)

func ParseFlagIsDebugLogger() bool {
	flag.Parse()
	if !*IsDebugLoggerFlag {
		return *IsDebugLoggerShortFlag
	}
	return *IsDebugLoggerFlag
}
