package post

import (
	"encoding/json"
	"log"

	"github.com/jinzhu/copier"
	"github.com/programzheng/base/pkg/function"
	"github.com/programzheng/base/pkg/model/post"
	"github.com/programzheng/base/pkg/service/file"
)

type Post struct {
	ID      uint        `json:"id"`
	Title   string      `json:"title"`
	Summary string      `json:"summary"`
	Detail  string      `json:"detail"`
	Files   interface{} `json:"files"`

	PageNum  int `form:"page_num" json:"page_num"`   //頁數*筆數,從0(代表第一頁)開始
	PageSize int `form:"page_size" json:"page_size"` //從PageNum之後取出的筆數
}

type Posts []Post

var (
	module string = "posts"
)

func (p *Post) Add() (Post, error) {
	fileReferenceJSON, err := json.Marshal(p.Files)
	if err != nil {
		log.Fatal("add "+module+" error", err)
	}
	fileReferenceJSONString := string(fileReferenceJSON)
	fileReference := function.CreateSHA1(fileReferenceJSONString)
	modelPost := post.Post{
		Title:         p.Title,
		Summary:       p.Summary,
		Detail:        p.Detail,
		FileReference: fileReference,
	}

	result, err := modelPost.Add()
	if err != nil {
		return Post{}, err
	}

	post := Post{}

	copier.Copy(&post, &result)

	batchUpdates := make(map[string]interface{})
	batchUpdates["reference"] = fileReference
	files, err := file.BatchUpdates(p.Files, nil, batchUpdates)
	if err != nil {
		return Post{}, err
	}
	post.Files = files
	return post, nil
}

func (p *Post) Get() ([]Post, error) {
	modelPosts, err := post.Get(p.PageNum, p.PageSize, p.getMaps())
	if err != nil {
		return nil, err
	}
	servicePosts := Posts{}
	copier.Copy(&servicePosts, &modelPosts)
	return servicePosts, nil
}

func (p *Post) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	maps["deleted_at"] = nil
	return maps
}
