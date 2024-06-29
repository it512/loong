package task

import (
	"context"

	"github.com/it512/loong"
)

type ReassignCmd struct {
	TaskID string

	*loong.Engine
}

func (r *ReassignCmd) Init(ctx context.Context, e *loong.Engine) error {
	r.Engine = e
	return nil
}

func (r ReassignCmd) Do(ctx context.Context) error {
	u := &loong.UserTask{}
	if err := r.Engine.Storer.LoadUserTask(ctx, r.TaskID, u); err != nil {
		return err
	}
	return nil
}
