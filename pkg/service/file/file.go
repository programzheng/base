package file

import (
	"github.com/jinzhu/copier"
	"github.com/programzheng/base/pkg/model/file"
)

type File struct {
	ID     uint   `json:"id"`
	System string `json:"-"`
	Type   string `json:"type"`
	Path   string `json:"path"`
	Name   string `json:"-"`
}

type Files []File

func (f *File) Add() (File, error) {
	model := file.File{
		System: f.System,
		Type:   f.Type,
		Path:   f.Path,
		Name:   f.Name,
	}
	modelFile, err := model.Add()
	if err != nil {
		return File{}, err
	}
	serviceFile := File{}
	copier.Copy(&serviceFile, &modelFile)
	return serviceFile, nil
}

func Get(ids []interface{}) (Files, error) {
	modelFiles, err := file.Get(ids, getMaps())
	if err != nil {
		return nil, err
	}
	serviceFiles := Files{}
	copier.Copy(&serviceFiles, &modelFiles)
	// for _, serviceFile := range serviceFiles {
	// 	serviceFile.Path = getResponseFilePath() + "/" + serviceFile.Path + serviceFile.Name
	// }
	return serviceFiles, nil
}

func BatchUpdates(ids interface{}, maps interface{}, updates interface{}) (Files, error) {
	modelFiles, err := file.BatchUpdates(ids, getMaps(), updates)
	if err != nil {
		return nil, err
	}
	serviceFiles := Files{}
	copier.Copy(&serviceFiles, &modelFiles)
	// for _, serviceFile := range serviceFiles {
	// 	serviceFile.Path = getResponseFilePath() + "/" + serviceFile.Path + serviceFile.Name
	// }
	return serviceFiles, nil
}

func getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	maps["deleted_at"] = nil
	return maps
}
