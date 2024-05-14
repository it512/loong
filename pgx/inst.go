package pgx

import (
	"context"

	"github.com/it512/loong"
	"github.com/it512/loong/pgx/internal/ent/inst"
)

func (m *Store) LoadProcInst(ctx context.Context, instID string, p *loong.ProcInst) error {
	q := m.Inst.Query()
	q.Where(inst.InstIDEQ(instID), inst.StatusEQ(1))
	inst, err := q.Only(ctx)
	if err != nil {
		return err
	}

	p.InstID = inst.InstID
	p.ProcID = inst.ProcID
	p.BusiKey = inst.BusiKey
	p.BusiType = inst.BusiType
	p.Starter = inst.Starter
	p.StartTime = inst.StartTime
	p.Status = inst.Status
	p.Init = inst.Init

	return nil
}

func (m *Store) CreateProcInst(ctx context.Context, procInst *loong.ProcInst) error {
	return m.Inst.Create().
		SetInstID(procInst.InstID).
		SetProcID(procInst.ProcID).
		SetBusiKey(procInst.BusiKey).
		SetBusiType(procInst.BusiType).
		SetStarter(procInst.Starter).
		SetStatus(procInst.Status).
		SetInit(procInst.Init).
		SetStartTime(procInst.StartTime).
		Exec(ctx)
}

func (m *Store) EndProcInst(ctx context.Context, procInst *loong.ProcInst) error {
	return m.Inst.Update().Where(inst.InstIDEQ(procInst.InstID)).
		SetStatus(procInst.Status).
		SetEndTime(procInst.EndTime).
		Exec(ctx)
}
