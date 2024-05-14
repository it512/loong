package gql

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/it512/loong"
	"github.com/it512/loong/todo"
	"github.com/it512/loong/web/gql/graph"
)

func New(engine *loong.Engine, todo todo.Todo, exts ...graphql.HandlerExtension) http.Handler {
	s := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{
		Resolvers: &graph.Resolver{Engine: engine, Todo: todo},
	}))
	for _, ext := range exts {
		s.Use(ext)
	}
	return s
}

func Playground(title, url string) http.Handler {
	return playground.ApolloSandboxHandler(title, url)
}
