package helper

import (
	"os"
	"runtime"
	"strings"
)

func GetOS() string {
	return runtime.GOOS
}

func GetPathSeparator() string {
	var separator string
	switch GetOS() {
	case "windows":
		separator = "\\"
	case "linux":
	default:
		separator = "/"
	}
	return separator
}

func IsRunnigTest() bool {
	if strings.HasSuffix(os.Args[0], ".test") {
		return true
	}
	return false
}
