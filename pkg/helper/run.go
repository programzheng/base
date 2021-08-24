package helper

import (
	"fmt"
	"os"
	"runtime"
	"strings"
)

func GetOS() string {
	fmt.Println(runtime.GOOS)
	return runtime.GOOS
}

func GetPathSeparator() string {
	var separator string
	switch GetOS() {
	case "windows":
		separator = "\\"
	case "linux":
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
