package helper

import (
	"mime"
	"os"
)

func CheckDirectoryExists(dir string) bool {
	_, err := os.Stat(dir)
	return !os.IsNotExist(err)
}

func GetFileContentType(extension string) string {
	return mime.TypeByExtension(extension)
}
