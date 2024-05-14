package graph

import (
	"github.com/it512/loong"
	"github.com/it512/loong/todo"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	*loong.Engine
	todo.Todo
}
