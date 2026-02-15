package flags

import "flag"

var (
	IsLogsInFile = flag.Bool("writingLogs", false, "flag for writing logs in file")
)

func ParseFlagIsLogsInFile() bool {
	flag.Parse()
	return *IsLogsInFile
}
