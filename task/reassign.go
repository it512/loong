package task

import (
	"context"
	"time"

	"github.com/it512/loong"
)

type ReassignCmd struct {
	TaskID string

	*loong.Engine

	ReassingFn func(context.Context, *loong.UserTask) error
}

func (r *ReassignCmd) Bind(ctx context.Context, e *loong.Engine) error {
	r.Engine = e
	return nil
}

func (r ReassignCmd) Do(ctx context.Context) error {
	u := loong.UserTask{}
	if err := r.Engine.Storer.LoadUserTask(ctx, r.TaskID, &u); err != nil {
		return err
	}

	ucopy := u // copy

	if err := r.ReassingFn(ctx, &ucopy); err != nil {
		return err
	}

	if err := r.Storer.EndUserTask(ctx, u); err != nil {
		return err
	}

	u.TaskID = r.Engine.NewID()
	u.Assignee = ucopy.Assignee
	u.CandidateGroups = ucopy.CandidateGroups
	u.CandidateUsers = ucopy.CandidateUsers
	u.StartTime = time.Now()

	return r.Engine.Storer.CreateTasks(ctx, u)
}
