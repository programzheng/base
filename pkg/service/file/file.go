package file

import (
	"github.com/programzheng/base/pkg/model/file"
)

type File struct {
	System string
	Type   string
	Path   string
	Name   string
}

func (f *File) Add() (uint, error) {
	model := file.File{
		System: f.System,
		Type:   f.Type,
		Path:   f.Path,
		Name:   f.Name,
	}
	ID, err := file.Add(model)
	if err != nil {
		return 0, err
	}
	return ID, nil
}
