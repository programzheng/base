package post

import (
	"github.com/programzheng/base/pkg/model/post"
	"github.com/programzheng/base/pkg/service"
	"github.com/programzheng/base/pkg/service/file"

	"github.com/jinzhu/copier"
)

type Post struct {
	ID      uint            `json:"id"`
	Title   string          `json:"title"`
	Summary string          `json:"summary"`
	Detail  string          `json:"detail"`
	Files   []file.LinkFile `json:"files"`
}

func (p *Post) Add() (*Post, error) {
	fileHashIds, fileReference := file.GetHashIdsAndReferenceByBase64LinkFiles(p.Files)

	modelPost := post.Post{
		Title:         p.Title,
		Summary:       p.Summary,
		Detail:        p.Detail,
		FileReference: fileReference,
	}

	result, err := modelPost.Add()
	if err != nil {
		return nil, err
	}

	batchUpdates := make(map[string]interface{}, 1)
	batchUpdates["reference"] = fileReference
	_, err = file.BatchUpdatesByHashIDs(*fileHashIds, nil, batchUpdates)
	if err != nil {
		return nil, err
	}

	post := Post{}

	copier.Copy(&post, &result)
	post.Files = file.GetLinkFilesByReference(modelPost.FileReference)

	return &post, nil
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
		servicePosts[index].Files = file.GetLinkFilesByReference(modelPost.FileReference)
	}

	return servicePosts, nil
}

func (p *Post) FindById() (*Post, error) {
	modelPost, err := post.FindById(p.ID)
	if err != nil {
		return nil, err
	}
	var servicePost *Post
	copier.Copy(&servicePost, &modelPost)

	return servicePost, nil
}

func (p *Post) UpdateByID(id uint) (*Post, error) {
	modelPost, err := post.FindById(id)
	if err != nil {
		return nil, err
	}

	fileHashIds, _ := file.GetHashIdsAndReferenceByBase64LinkFiles(p.Files)

	updates := map[string]interface{}{
		"title":   p.Title,
		"Summary": p.Summary,
		"Detail":  p.Detail,
	}
	result, err := post.UpdateByModel(*modelPost, updates)
	if err != nil {
		return nil, err
	}

	batchUpdates := make(map[string]interface{}, 1)
	batchUpdates["reference"] = result.FileReference
	_, err = file.BatchUpdatesByHashIDs(*fileHashIds, nil, batchUpdates)
	if err != nil {
		return nil, err
	}

	post := Post{}
	copier.Copy(&post, &result)
	post.Files = file.GetLinkFilesByReference(modelPost.FileReference)

	return &post, nil
}

func (p *Post) DelByID(id uint) error {
	err := post.DelByID(id)
	if err != nil {
		return err
	}
	return nil
}

func (p *Post) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	maps["deleted_at"] = nil
	return maps
}
