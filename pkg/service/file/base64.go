package file

import (
	"context"
	"encoding/base64"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/programzheng/base/pkg/filesystem"
	"github.com/programzheng/base/pkg/helper"
)

type Base64File struct {
	ContentType string
	Extension   string
	Encode      string
}

func CheckMatchBase64Image(originB64 string) bool {
	var re = regexp.MustCompile(`^data:image\/(?:gif|png|jpeg|bmp|webp|svg\+xml)(?:;charset=utf-8)?;base64,(?:[A-Za-z0-9]|[+/])+={0,2}`)

	return re.MatchString(originB64)
}

func ConvertBase64ToBase64File(originB64 string) *Base64File {
	var re = regexp.MustCompile(`^data:|;base64,`)

	dec := re.Split(originB64, 3)

	contentType := dec[1]
	ext := helper.GetFileExtensionByContentType(contentType)

	b64File := &Base64File{
		ContentType: contentType,
		Extension:   ext,
		Encode:      dec[2],
	}

	return b64File
}

func (b64File *Base64File) Base64FileConvertToFile(filePath string, fileFullName string) *os.File {
	dec, err := base64.StdEncoding.DecodeString(b64File.Encode)

	if err != nil {
		log.Panic(err)
	}

	f, err := os.CreateTemp("", "")
	if err != nil {
		log.Panic(err)
	}

	if _, err := f.Write(dec); err != nil {
		log.Panic(err)
	}
	if err := f.Sync(); err != nil {
		log.Panic(err)
	}

	// go to begginng of file
	f.Seek(0, 0)

	// output file
	return f
}

func AddFileByBase64(b64s []string) ([]string, *string) {
	ctx, cancel := context.WithCancel(context.Background())
	fileHashIds := make([]string, len(b64s))
	for b64Index, b64 := range b64s {
		if !CheckMatchBase64Image(b64) {
			continue
		}
		b64File := ConvertBase64ToBase64File(b64)
		fileFullName := helper.CreateUuid() + b64File.Extension
		uf := b64File.Base64FileConvertToFile("", fileFullName)
		staticFile := filesystem.Create("").Upload(ctx, fileFullName, uf)
		if staticFile == nil {
			log.Fatal("add by file error")
		}
		fileService := File{
			System:      staticFile.System,
			Type:        staticFile.Type,
			Path:        staticFile.Path,
			Name:        staticFile.Name,
			ThirdPatyID: staticFile.ThirdPatyID,
		}
		file, err := fileService.Add()
		if err != nil {
			log.Fatalf("add by file error:%v", err)
		}
		fileHashIds[b64Index] = file.HashID
		defer os.Remove(uf.Name())
	}
	cancel()

	reference := helper.CreateMD5(strings.Join(fileHashIds, ","))

	return fileHashIds, &reference
}
