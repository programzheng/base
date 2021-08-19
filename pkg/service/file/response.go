package file

import (
	"base/pkg/filesystem"
)

func getResponseFilePath() string {
	return filesystem.Driver.GetHostURL()
}
