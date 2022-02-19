package helper

import (
	"mime"
	"os"
	"regexp"

	log "github.com/sirupsen/logrus"
)

func CheckDirectoryExists(dir string) bool {
	_, err := os.Stat(dir)
	return !os.IsNotExist(err)
}

func GetFileContentType(extension string) string {
	return mime.TypeByExtension(extension)
}

func GetFileExtensionByContentType(contentType string) string {
	exts, err := mime.ExtensionsByType(contentType)
	if err != nil {
		log.Fatal("helper GetFileExtensionByContentType func error:", err)
	}
	var re = regexp.MustCompile(`/(.*)$`)
	t := re.FindStringSubmatch(contentType)[1]
	result := exts[0]
	for _, ext := range exts {
		if "."+t == ext {
			result = ext
			break
		}
	}

	return result
}
