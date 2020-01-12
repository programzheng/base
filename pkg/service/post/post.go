package post

import (
	"log"

	"github.com/programzheng/base/pkg/function"
	"github.com/programzheng/base/pkg/model/post"
	"github.com/programzheng/base/pkg/service/file"
)

type Post struct {
	ID      uint          `json:"id"`
	Title   string        `json:"title"`
	Summary string        `json:"summary"`
	Detail  string        `json:"detail"`
	Files   []interface{} `json:"files"`

	PageNum  int `form:"page_num"`  //頁數*筆數,從0(代表第一頁)開始
	PageSize int `form:"page_size"` //從PageNum之後取出的筆數
}

var (
	module string = "posts"
)

func (p *Post) Add() error {
	files := function.ConvertInterfaceToIntMap(p.Files)
	fileReferenceByte, err := function.GetBytes(files)
	if err != nil {
		log.Fatal("add "+module+" error", err)
	}
	fileReference := function.CreateSHA1(fileReferenceByte)
	model := post.Post{
		Title:         p.Title,
		Summary:       p.Summary,
		Detail:        p.Detail,
		FileReference: fileReference,
	}
	result, err := model.Add()
	if err != nil {
		return err
	}
	function.GetJSON(result)
	return nil
}

func (p *Post) Get() ([]Post, error) {
	modelPosts, err := post.Get(p.PageNum, p.PageSize, p.getMaps())
	if err != nil {
		return nil, err
	}
	servicePosts := []Post{}
	for _, modelPost := range modelPosts {
		// TODO: response file 類型錯誤
		files := file.NewResponseFiles(modelPost.Files)
		servicePost := Post{
			ID:      modelPost.ID,
			Title:   modelPost.Title,
			Summary: modelPost.Summary,
			Detail:  modelPost.Detail,
			Files:   files,
		}
		servicePosts = append(servicePosts, servicePost)
	}
	return servicePosts, nil
}

func (p *Post) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	maps["deleted_at"] = nil
	return maps
}
