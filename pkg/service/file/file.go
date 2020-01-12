package file

import (
	"github.com/programzheng/base/pkg/model"
	"github.com/programzheng/base/pkg/model/file"
)

type File struct {
	System string
	Type   string
	Path   string
	Name   string
}

func (f *File) Add() (ResponseFile, error) {
	model := file.File{
		System: f.System,
		Type:   f.Type,
		Path:   f.Path,
		Name:   f.Name,
	}
	result, err := model.Add()
	if err != nil {
		return ResponseFile{}, err
	}
	responseFile := NewResponseFile(result)
	return *responseFile, nil
}

func Get() ([]*file.File, error) {
	files, err := file.Get(getMaps())
	if err != nil {
		return nil, err
	}
	return files, nil
}

func GetByID(ids model.JSON) ([]*file.File, error) {
	files, err := file.Get(getMapsByID(ids))
	if err != nil {
		return nil, err
	}
	return files, nil
}

func getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	maps["deleted_at"] = nil
	return maps
}

func getMapsByID(ids model.JSON) map[string]interface{} {
	maps := make(map[string]interface{})
	maps["id"] = ids
	maps["deleted_at"] = nil
	return maps
}
