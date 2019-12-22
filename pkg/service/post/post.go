package post

import (
	"github.com/programzheng/base/pkg/model/post"
)

type Post struct {
	Title   string `json:"title"`
	Summary string `json:"summary"`
	// Detail  string `sql:"type:text"`
	// Images  string `sql:"type:json`

	PageNum  int `form:"page_num"`  //頁數*筆數,從0(代表第一頁)開始
	PageSize int `form:"page_size"` //從PageNum之後取出的筆數
}

func (p *Post) Add() error {
	model := post.Post{
		Title:   p.Title,
		Summary: p.Summary,
	}
	if err := post.Add(model); err != nil {
		return err
	}
	return nil
}

func (p *Post) Get() ([]*post.Post, error) {
	posts, err := post.Get(p.PageNum, p.PageSize, p.getMaps())
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (p *Post) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	maps["deleted_at"] = nil
	return maps
}
