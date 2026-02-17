package flags

import "flag"

var (
	IsLogsInFileFlag      = flag.Bool("writingLogs", false, "flag for writing logs in file")
	IsLogsInFileShortFlag = flag.Bool("l", false, "flag for writing logs in file short")
)

func ParseFlagIsLogsInFile() bool {
	flag.Parse()
	if !*IsLogsInFileFlag {
		return *IsLogsInFileShortFlag
	}
	return *IsLogsInFileFlag
}
