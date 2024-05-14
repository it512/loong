package loong

import (
	"context"
	"time"
)

type Out interface {
	GetOutgoingAssociation() []string
}

type ProcInst struct {
	InstID string
	ProcID string

	Starter  string
	Operator string

	BusiKey  string
	BusiType string

	/*
	EndCode string
	EndName string
	EndType string
	*/

	StartTime time.Time
	EndTime   time.Time

	Status int

	Init Var

	Template *Template
	*Engine
}

type Exec struct {
	ExecID string

	ParentForkID string // 上一级fork , 顶级为InstID
	ForkID       string // 当前fork

	ForkTag string // 谁fork的
	JoinTag string // 谁join的
	OutTag  string // fork的出口
	InTag   string // join的入口

	GwType   int // 网关类型 // 并行，包容
	ForkMode int // fork 模式

	Status int

	Input Var

	*ProcInst
}

func (e Exec) Eval(ctx context.Context, el string) (any, error) {
	return e.Evaluator.Eval(ctx, el, e)
}

func (c Exec) Do(ctx context.Context) error {
	return nil
}

func (c Exec) Emit(ctx context.Context, emt Emitter) error {
	return nil
}

func (c Exec) Type() ActivityType {
	return NotApplicable
}

func (e Exec) EmitDefault(ctx context.Context, o Out, emt Emitter) error {
	flows := e.ProcInst.Template.FindSequenceFlows(o.GetOutgoingAssociation())
	out, err := choose(ctx, e, flows)
	if err != nil {
		return err
	}
	f := chooseDefault(o, out)
	return emt.Emit(&sequenceFlow{TSequenceFlow: f, Exec: e})
}

func (e Exec) isTop() bool {
	return e.ParentForkID == ""
}

func (e Exec) parent() Exec {
	return Exec{
		ForkID:   e.ParentForkID,
		ForkTag:  e.ForkTag,
		GwType:   e.GwType,
		ForkMode: e.ForkMode,
		Status:   STATUS_START,
		ProcInst: e.ProcInst,
	}
}

func (e Exec) forkOut(out []string) (forkID string, x []Exec) {
	forkID = e.Engine.NewID()
	for _, o := range out {
		x = append(x,
			Exec{
				ExecID:       e.Engine.NewID(),
				ForkID:       forkID,
				ParentForkID: e.ForkID,
				ForkTag:      e.ForkTag,
				OutTag:       o,
				GwType:       e.GwType,
				ForkMode:     e.ForkMode,
				Status:       STATUS_START,
				ProcInst:     e.ProcInst,
			})
	}

	return
}

func (e Exec) top() Exec {
	return Exec{
		ForkTag:  e.ForkTag,
		Status:   STATUS_START,
		ProcInst: e.ProcInst,
	}
}

func (e Exec) children(out []string) (x []Exec) {
	for _, o := range out {
		x = append(x,
			Exec{
				ExecID:       e.Engine.NewID(),
				ForkID:       e.ForkID,
				ParentForkID: e.ParentForkID,
				ForkTag:      e.ForkTag,
				OutTag:       o,
				GwType:       e.GwType,
				ForkMode:     e.ForkMode,
				Status:       STATUS_START,
				ProcInst:     e.ProcInst,
			})
	}
	return
}

func (e Exec) empty() Exec {
	return Exec{
		Status:   STATUS_START,
		ProcInst: e.ProcInst,
	}
}
