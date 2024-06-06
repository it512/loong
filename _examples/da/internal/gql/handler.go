package gql

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/it512/da/internal/gql/graph"
	"github.com/it512/loong"
)

func New(engine *loong.Engine, exts ...graphql.HandlerExtension) http.Handler {
	s := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{
		Resolvers: &graph.Resolver{Engine: engine},
	}))
	for _, ext := range exts {
		s.Use(ext)
	}
	return s
}

func Playground(title, url string) http.Handler {
	return playground.ApolloSandboxHandler(title, url)
}
