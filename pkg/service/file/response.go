package file

import (
	"strconv"

	"github.com/programzheng/base/pkg/filesystem"
	"github.com/programzheng/base/pkg/model/file"
)

type ResponseFile struct {
	ID   uint
	Path string
}

func NewResponseFiles(files []file.File) *[]ResponseFile {
	responseFiles := new([]ResponseFile{})
	for _, file := range files {
		responseFile := ResponseFile{
			ID:   file.ID,
			Path: getResponseFilePath() + strconv.FormatUint(uint64(file.ID), 10),
		}
		responseFiles = append(responseFiles, responseFile)
	}
	return responseFiles
}

func NewResponseFile(file file.File) *ResponseFile {

	responseFile := ResponseFile{
		ID:   file.ID,
		Path: getResponseFilePath() + strconv.FormatUint(uint64(file.ID), 10),
	}

	return &responseFile
}

func getResponseFilePath() string {
	return filesystem.Driver.GetHostURL()
}
