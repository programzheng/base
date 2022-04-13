package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/programzheng/base/graph/generated"
	"github.com/programzheng/base/graph/model"
	"github.com/programzheng/base/pkg/helper"
	"github.com/programzheng/base/pkg/service"
	"github.com/programzheng/base/pkg/service/admin"
	"github.com/programzheng/base/pkg/service/post"
)

func (r *authAdminOpsResolver) Login(ctx context.Context, obj *model.AuthAdminOps, account string, password string) (interface{}, error) {
	token, err := admin.Login("", account, password)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (r *authAdminOpsResolver) Register(ctx context.Context, obj *model.AuthAdminOps, input model.RegisterAdmin) (interface{}, error) {
	adminService := admin.Admin{
		Account:  input.Account,
		Password: input.Password,
		Profile: admin.AdminProfile{
			Name: input.Name,
		},
	}
	//hash password
	adminService.Password = helper.CreateHash(adminService.Password)
	result, err := adminService.Add()
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *mutationResolver) AuthAdmin(ctx context.Context) (*model.AuthAdminOps, error) {
	return &model.AuthAdminOps{}, nil
}

func (r *queryResolver) AuthAdmin(ctx context.Context) (interface{}, error) {
	return map[string]interface{}{
		"code":  200,
		"value": nil,
	}, nil
}

func (r *queryResolver) Posts(ctx context.Context) ([]*post.Post, error) {
	var postService post.Post
	page := service.Page{
		Num:  1,
		Size: 10,
	}
	posts, err := postService.Get(page)
	if err != nil {
		return nil, err
	}
	result := []*post.Post{}
	for i := range posts {
		result = append(result, &posts[i])
	}
	return result, nil
}

// AuthAdminOps returns generated.AuthAdminOpsResolver implementation.
func (r *Resolver) AuthAdminOps() generated.AuthAdminOpsResolver { return &authAdminOpsResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type authAdminOpsResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
