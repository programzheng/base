package post

import (
	"log"

	"github.com/programzheng/base/pkg/model/post"
	"github.com/programzheng/base/pkg/service"
	"github.com/programzheng/base/pkg/service/file"

	"github.com/jinzhu/copier"
)

type Post struct {
	ID      uint     `json:"id"`
	Title   string   `json:"title"`
	Summary string   `json:"summary"`
	Detail  string   `json:"detail"`
	Files   []string `json:"files"`
}

var (
	module string = "posts"
)

func (p *Post) Add() (Post, error) {
	fileHashIds := []string{}
	fileReference := ""
	if len(p.Files) > 0 {
		fileHashIds, fileReference = file.AddFileByBase64(p.Files)
	}

	modelPost := post.Post{
		Title:         p.Title,
		Summary:       p.Summary,
		Detail:        p.Detail,
		FileReference: &fileReference,
	}

	result, err := modelPost.Add()
	if err != nil {
		return Post{}, err
	}

	batchUpdates := make(map[string]interface{}, 1)
	batchUpdates["reference"] = fileReference
	_, err = file.BatchUpdatesByHashIDs(fileHashIds, func() map[string]interface{} {
		maps := make(map[string]interface{})
		return maps
	}, batchUpdates)
	if err != nil {
		log.Fatalf("BatchUpdatesByHashIDs error:%v", err)
	}

	post := Post{}

	copier.Copy(&post, &result)
	post.Files = file.GetFileOpenLinksByReference(*modelPost.FileReference)

	return post, nil
}
func (p *Post) GetTotalNumber() (int64, error) {
	count, err := post.GetTotalNumber(p.getMaps())
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (p *Post) Get(page service.Page) ([]Post, error) {
	modelPosts, err := post.Get(page.GetSqlOffset(), page.GetSqlLimit(), p.getMaps())
	if err != nil {
		return nil, err
	}
	servicePosts := make([]Post, len(modelPosts))
	copier.Copy(&servicePosts, &modelPosts)
	for index, modelPost := range modelPosts {
		servicePosts[index].Files = file.GetFileOpenLinksByReference(*modelPost.FileReference)
	}

	return servicePosts, nil
}

func (p *Post) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	maps["deleted_at"] = nil
	return maps
}
