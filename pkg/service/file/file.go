package file

import (
	"github.com/programzheng/base/pkg/model"
	"github.com/programzheng/base/pkg/model/file"
	"github.com/spf13/viper"

	"github.com/jinzhu/copier"
)

type File struct {
	ID          uint `json:"-"`
	HashID      string
	Reference   *string `json:"-"`
	System      string  `json:"-"`
	Type        string
	Path        string `json:"-"`
	Name        string `json:"-"`
	ThirdPatyID string `json:"-"`
}

type Files []File

func (f *File) Add() (File, error) {
	model := file.File{
		Reference:   f.Reference,
		System:      f.System,
		Type:        f.Type,
		Path:        f.Path,
		Name:        f.Name,
		ThirdPatyID: f.ThirdPatyID,
	}
	modelFile, err := model.Add()
	if err != nil {
		return File{}, err
	}
	serviceFile := File{}
	copier.Copy(&serviceFile, &modelFile)
	return serviceFile, nil
}

func Get(ids []interface{}, fn func() map[string]interface{}) (Files, error) {
	maps := fn()
	modelFiles, err := file.Get(ids, getMaps(maps))
	if err != nil {
		return nil, err
	}
	serviceFiles := Files{}
	copier.Copy(&serviceFiles, &modelFiles)

	return serviceFiles, nil
}

func GetFileOpenLinksByReference(reference *string) []string {
	serviceFiles, err := Get(nil, func() map[string]interface{} {
		maps := make(map[string]interface{}, 1)
		maps["reference"] = reference
		return maps
	})
	if err != nil {
		return nil
	}

	links := make([]string, len(serviceFiles))
	for index, serviceFile := range serviceFiles {
		links[index] = serviceFile.GetOpenLink()
	}

	return links
}

func BatchUpdates(fn func() map[string]interface{}, updates interface{}) (Files, error) {
	maps := fn()

	var modelFiles []file.File

	err := model.GetDB().Model(&modelFiles).Where(maps).Updates(updates).Find(&modelFiles).Error
	if err != nil {
		return nil, err
	}

	serviceFiles := Files{}
	copier.Copy(&serviceFiles, &modelFiles)

	return serviceFiles, nil
}

func BatchUpdatesByHashIDs(hashIDs []string, fn func() map[string]interface{}, updates map[string]interface{}) (Files, error) {
	maps := fn()

	var modelFiles []file.File

	err := model.GetDB().Model(&modelFiles).Select("reference").Where("hash_id IN (?)", hashIDs).Where(maps).Updates(updates).Find(&modelFiles).Error
	if err != nil {
		return nil, err
	}
	serviceFiles := Files{}
	copier.Copy(&serviceFiles, &modelFiles)

	return serviceFiles, nil
}

func getMaps(maps map[string]interface{}) map[string]interface{} {
	maps["deleted_at"] = nil
	return maps
}

func (f *File) GetOpenLink() string {
	link := ""
	switch f.System {
	case "local":
		link = "//" + viper.Get("APP_URL").(string) + ":" + viper.Get("APP_PORT").(string) + "/files/" + f.HashID
	case "cloudinary":
		link = f.Path
	}
	return link
}

func (f *File) GetLocalLink() string {
	link := ""
	switch f.System {
	case "local":
		link = f.Path + "/" + f.Name
	case "cloudinary":
		link = f.Path
	}
	return link
}
