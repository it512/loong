package loong

import (
	"context"
	"time"
)

type ProcInst struct {
	InstID string `json:"inst_id,omitempty"`
	ProcID string `json:"proc_id,omitempty"`

	Starter string `json:"starter,omitempty"`

	BusiKey  string `json:"busi_key,omitempty"`
	BusiType string `json:"busi_type,omitempty"`

	/*
		EndCode string
		EndName string
		EndType string
	*/

	StartTime time.Time `json:"start_time,omitempty"`
	EndTime   time.Time `json:"end_time,omitempty"`

	Status int `json:"status,omitempty"`

	Init Var `json:"init,omitempty"`

	*Template `json:"-"`
	*Engine   `json:"-"`
}

type Exec struct {
	ExecID string `json:"exec_id,omitempty"`

	ParentForkID string `json:"parent_fork_id,omitempty"` // 上一级fork , 顶级为InstID
	ForkID       string `json:"fork_id,omitempty"`        // 当前fork

	ForkTag string `json:"fork_tag,omitempty"` // 谁fork的
	JoinTag string `json:"join_tag,omitempty"` // 谁join的
	OutTag  string `json:"out_tag,omitempty"`  // fork的出口
	InTag   string `json:"in_tag,omitempty"`   // join的入口

	GwType   int `json:"gw_type,omitempty"`   // 网关类型 // 并行，包容
	ForkMode int `json:"fork_mode,omitempty"` // fork 模式

	Status int `json:"status,omitempty"`

	Input Var `json:"input,omitempty"`

	*ProcInst

	elementID string `json:"-"` // 当前bpmnElementID
}

func (e Exec) Eval(ctx context.Context, el string) (any, error) {
	return e.Evaluator.Eval(ctx, el, e)
}

func (s Exec) GetTaskDefinition(ctx context.Context) TaskDefinition {
	return &taskDef{id: s.elementID}
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
