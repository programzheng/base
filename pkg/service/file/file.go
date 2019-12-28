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

func (f *File) Add() error {
	model := file.File{
		System: f.System,
		Type:   f.Type,
		Path:   f.Path,
		Name:   f.Name,
	}
	if err := file.Add(model); err != nil {
		return err
	}
	return nil
}
