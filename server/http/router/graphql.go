package router

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/programzheng/base/graph"
	"github.com/programzheng/base/graph/generated"
	"github.com/programzheng/base/pkg/directive"
	"github.com/programzheng/base/pkg/middleware"
)

func setGraphqlRouter(router *gin.Engine) {
	graphqlGroup := router.Group("/graphql")
	{
		graphqlGroup.GET("", playgroundHandler())
	}
	graphqlGroup.Use(middleware.GraphqlValidJSONWebToken())
	{
		graphqlGroup.POST("", graphqlHandler())
	}
}

// Defining the Graphql handler
func graphqlHandler() gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	c := generated.Config{Resolvers: &graph.Resolver{}}
	c.Directives.AuthAdmin = directive.AuthAdmin
	h := handler.NewDefaultServer(generated.NewExecutableSchema(c))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/graphql")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
