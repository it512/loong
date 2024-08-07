package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"

	"github.com/it512/da/internal/gql/graph/model"
	"github.com/it512/loong"
)

// StartProc is the resolver for the startProc field.
func (r *mutationResolver) StartProc(ctx context.Context, input model.StartProcCmd) (*model.ProcReturn, error) {
	err := r.Engine.RunActivityCmd(ctx, &loong.StartProcCmd{
		ProcID:   input.ProcID,
		Starter:  input.Starter,
		BusiKey:  input.BusiKey,
		BusiType: input.BusiType,
		Input:    input.Input,
		Var:      input.Var,
		Tags:     input.Tags,
	})

	return &model.ProcReturn{InstID: "!"}, err
}

// CommitTask is the resolver for the commitTask field.
func (r *mutationResolver) CommitTask(ctx context.Context, input model.UserTaskCommitCmd) (*model.CommitTaskReturn, error) {
	err := r.Engine.RunActivityCmd(ctx, &loong.UserTaskCommitCmd{
		TaskID:   input.TaskID,
		InstID:   input.InstID,
		Operator: input.Operator,
		Input:    input.Input,
		Var:      input.Var,
		Result:   input.Result,
		Version:  input.Version,
	})

	return &model.CommitTaskReturn{TaskID: input.TaskID}, err
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
