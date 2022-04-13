package directive

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"

	"github.com/programzheng/base/pkg/middleware"
)

func AuthAdmin(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	tokenData := middleware.CtxValue(ctx)
	if tokenData == "" {
		return nil, &gqlerror.Error{
			Message: "Access Denied",
		}
	}

	return next(ctx)
}
