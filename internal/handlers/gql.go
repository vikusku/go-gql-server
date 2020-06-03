package handlers

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/vikusku/go-gql-server/internal/gql"
	"github.com/vikusku/go-gql-server/internal/gql/resolvers"
)

func GraphqlHandler() gin.HandlerFunc {
	c := gql.Config{
		Resolvers: &resolvers.Resolver{},
	}

	h := handler.NewDefaultServer(gql.NewExecutableSchema(c))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func PlaygroundHandler(path string) gin.HandlerFunc {
	h := playground.Handler("Go GraphQL Server", path)
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
