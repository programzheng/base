package version

import (
	"fmt"
	"runtime"
)

var (
	Version string

	BuildData string
)

func PrintCLIVersion() string {
	return fmt.Sprintf(
		"version %s built on %s, %s",
		Version,
		BuildData,
		runtime.Version(),
	)
}
