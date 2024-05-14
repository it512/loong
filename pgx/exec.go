package pgx

import (
	"context"

	"github.com/it512/loong"
	"github.com/it512/loong/pgx/internal/ent"
	"github.com/it512/loong/pgx/internal/ent/exec"
)

// fork
func (m *Store) ForkExec(ctx context.Context, xs []loong.Exec) error {
	b := m.Exec.MapCreateBulk(xs, func(cr *ent.ExecCreate, i int) {
		cr.SetExecID(xs[i].ExecID).SetInstID(xs[i].InstID).SetForkID(xs[i].ForkID).
			SetParentForkID(xs[i].ParentForkID).
			SetOutTag(xs[i].OutTag).
			SetForkTag(xs[i].ForkTag).
			SetGwType(xs[i].GwType).
			SetStatus(xs[i].Status)
	})
	return b.Exec(ctx)
}

// join
func (m *Store) JoinExec(ctx context.Context, ex *loong.Exec) error {
	return m.Exec.Update().Where(exec.ExecIDEQ(ex.ExecID)).
		SetStatus(ex.Status).
		SetJoinTag(ex.JoinTag).
		Exec(ctx)
}

func (m *Store) LoadForks(ctx context.Context, forkID string) ([]loong.Exec, error) {
	xs, err := m.Exec.Query().Where(exec.ForkIDEQ(forkID)).All(ctx)
	if err != nil {
		return nil, err
	}
	var s []loong.Exec
	for _, x := range xs {
		s = append(s,
			loong.Exec{
				ExecID:       x.ExecID,
				ParentForkID: x.ParentForkID,
				ForkID:       x.ForkID,
				ForkTag:      x.ForkTag,
				JoinTag:      x.JoinTag,
				OutTag:       x.OutTag,
				GwType:       x.GwType,
				Status:       x.Status,
			})
	}
	return s, nil
}

func (m *Store) LoadExec(ctx context.Context, execID string, ex *loong.Exec) error {
	x, err := m.Exec.Query().Where(exec.ExecIDEQ(execID)).Only(ctx)
	if err != nil {
		return err
	}
	ex.ExecID = x.ExecID
	ex.ParentForkID = x.ParentForkID
	ex.ForkID = x.ForkID
	ex.ForkTag = x.ForkTag
	ex.JoinTag = x.JoinTag
	ex.OutTag = x.OutTag
	ex.GwType = x.GwType
	ex.Status = x.Status

	return nil
}
