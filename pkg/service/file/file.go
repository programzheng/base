package file

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strings"

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

type LinkFile struct {
	HashID string `json:"hash_id"`
	Type   string `json:"type"`
	Value  string `json:"value"`
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

func Get(where map[string]interface{}) (Files, error) {
	modelFiles, err := file.Get(getMaps(where))
	if err != nil {
		return nil, err
	}
	serviceFiles := Files{}
	copier.Copy(&serviceFiles, &modelFiles)

	return serviceFiles, nil
}

func GetHashIdsAndReferenceByBase64LinkFiles(lfs []LinkFile) ([]string, *string) {
	b64s := make([]string, 0, len(lfs))
	for _, lf := range lfs {
		//new file hashID is ""
		if lf.HashID != "" {
			continue
		}
		b64s = append(b64s, lf.Value)
	}

	fileHashIds, fileReference := AddFileByBase64(b64s)
	return fileHashIds, fileReference
}

func GetLinkFilesByReference(reference *string) []LinkFile {
	where := make(map[string]interface{}, 1)
	where["reference"] = reference
	serviceFiles, err := Get(where)
	if err != nil {
		return nil
	}

	lfs := make([]LinkFile, len(serviceFiles))
	for index, serviceFile := range serviceFiles {
		lfs[index].HashID = serviceFile.HashID
		lfs[index].Type = serviceFile.Type
		lfs[index].Value = serviceFile.GetOpenLink()
	}

	return lfs
}

func BatchUpdates(where map[string]interface{}, updates map[string]interface{}) (Files, error) {
	modelFiles, err := file.BatchUpdates(where, updates)
	if err != nil {
		return nil, err
	}
	serviceFiles := Files{}
	copier.Copy(&serviceFiles, &modelFiles)

	return serviceFiles, nil
}

func GetHashIDsByLinkFiles(lfs []LinkFile) []string {
	hashIDs := make([]string, len(lfs))
	for _, lf := range lfs {
		hashIDs = append(hashIDs, lf.HashID)
	}

	return hashIDs
}

func Delete(where map[string]interface{}) error {
	err := file.Delete(where)
	if err != nil {
		return err
	}
	return nil
}

func getMaps(maps map[string]interface{}) map[string]interface{} {
	maps["deleted_at"] = nil
	return maps
}

func ReplaceProtocolEmpty(path string) string {
	if strings.HasPrefix(path, "https:") {
		path = strings.Replace(path, "https:", "", 1)
	}
	if strings.HasPrefix(path, "http:") {
		path = strings.Replace(path, "http:", "", 1)
	}
	return path
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

func (f *File) GetBytes() ([]byte, error) {
	var bs []byte
	switch f.System {
	case "local":
		tmpBs, err := ioutil.ReadFile(f.GetLocalLink())
		if err != nil {
			return nil, err
		}
		bs = tmpBs
	case "cloudinary":
		response, err := http.Get(f.GetOpenLink())
		if err != nil {
			return nil, err
		}

		buf := new(bytes.Buffer)
		buf.ReadFrom(response.Body)
		reader := response.Body
		defer reader.Close()
		bs = buf.Bytes()
	}

	return bs, nil
}
