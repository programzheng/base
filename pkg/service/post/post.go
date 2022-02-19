package post

import (
	"github.com/programzheng/base/pkg/model/post"
	"github.com/programzheng/base/pkg/service/file"

	"github.com/jinzhu/copier"
)

type Post struct {
	ID      uint     `json:"id"`
	Title   string   `json:"title"`
	Summary string   `json:"summary"`
	Detail  string   `json:"detail"`
	Files   []string `json:"files"`

	PageNum  int `form:"page_num" json:"page_num"`   //頁數*筆數,從0(代表第一頁)開始
	PageSize int `form:"page_size" json:"page_size"` //從PageNum之後取出的筆數
}

var (
	module string = "posts"
)

func (p *Post) Add() (Post, error) {
	fileReference := ""
	if len(p.Files) > 0 {
		fileReference = file.AddFileByBase64(p.Files)
	}

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
	post.Files = file.GetFileOpenLinksByReference(modelPost.FileReference)

	return post, nil
}

func (p *Post) Get() ([]Post, error) {
	modelPosts, err := post.Get(p.PageNum, p.PageSize, p.getMaps())
	if err != nil {
		return nil, err
	}
	servicePosts := make([]Post, len(modelPosts))
	copier.Copy(&servicePosts, &modelPosts)
	for index, modelPost := range modelPosts {
		servicePosts[index].Files = file.GetFileOpenLinksByReference(modelPost.FileReference)
	}

	return servicePosts, nil
}

func (p *Post) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	maps["deleted_at"] = nil
	return maps
}
