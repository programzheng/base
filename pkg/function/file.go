package function

import (
	"log"
	"mime"
)

func GetFileContentType(extension string) string {
	mimeType := mime.TypeByExtension(extension)
	if mimeType == "" {
		log.Println("file upload mime type not found")
	}
	return mimeType
}
