package helper

import (
	"mime"
)

func GetFileContentType(extension string) string {
	return mime.TypeByExtension(extension)
}
