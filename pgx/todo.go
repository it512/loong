package pgx

import (
	"context"
	"slices"

	"github.com/it512/loong"
	"github.com/it512/loong/pgx/internal/ent"
	"github.com/it512/loong/pgx/internal/ent/usertask"
	"github.com/it512/loong/todo"
)

type TodoStore struct {
	*ent.Client
}

func todoPageDown(ctx context.Context, tp *todo.TodoQueryParam, client *ent.Client) (r todo.TodoPageResult, err error) {
	groups := slices.Concat(tp.CandidateGroups, tp.SudoGroups)

	q := client.UserTask.Query()
	q.Where(
		usertask.CandidateGroupsIn(groups...),
		usertask.TaskIDLT(tp.Last),
		usertask.StatusEQ(loong.STATUS_START),
	).
		Order(ent.Desc(usertask.FieldTaskID)).
		Limit(tp.Size)

	var ut []*ent.UserTask
	if ut, err = q.All(ctx); err != nil {
		return
	}

	for _, u := range ut {
		r.Items = append(r.Items, xx(u))
	}
	r.Count = len(r.Items)
	if r.Count > 0 {
		r.Max = r.Items[0].TaskID
		r.Mix = r.Items[r.Count-1].TaskID
	}

	return
}

func todoPageUp(ctx context.Context, tp *todo.TodoQueryParam, client *ent.Client) (r todo.TodoPageResult, err error) {
	groups := slices.Concat(tp.CandidateGroups, tp.SudoGroups)

	q := client.UserTask.Query()
	q.Where(
		usertask.CandidateGroupsIn(groups...),
		usertask.TaskIDGT(tp.Last),
		usertask.StatusEQ(loong.STATUS_START),
	).
		Order(ent.Desc(usertask.FieldTaskID)).
		Limit(tp.Size)

	var ut []*ent.UserTask
	if ut, err = q.All(ctx); err != nil {
		return
	}

	for _, u := range ut {
		r.Items = append(r.Items, xx(u))
	}

	r.Count = len(r.Items)
	if r.Count > 0 {
		r.Max = r.Items[0].TaskID
		r.Mix = r.Items[r.Count-1].TaskID
	}
	return
}

func xx(u *ent.UserTask) *todo.Item {
	return &todo.Item{
		TaskID:          u.TaskID,
		FormKey:         u.FormKey,
		BusiKey:         u.BusiKey,
		BusiType:        u.BusiType,
		ActName:         u.ActName,
		ActID:           u.ActID,
		Version:         u.Version,
		CandidateGroups: u.CandidateGroups,
	}
}

func (m *TodoStore) QueryTaskPage(ctx context.Context, tp *todo.TodoQueryParam) (todo.TodoPageResult, error) {
	if tp.Direction == 0 { //down
		return todoPageDown(ctx, tp, m.Client)
	}
	return todoPageUp(ctx, tp, m.Client)
}
