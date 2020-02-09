package file

import (
	"github.com/programzheng/base/pkg/filesystem"
)

func getResponseFilePath() string {
	return filesystem.Driver.GetHostURL()
}
