package loong

import (
	"context"
	"errors"
	"fmt"
	"maps"
	"time"

	"github.com/it512/loong/bpmn"
)

type StartProcCmd struct {
	ProcID   string         `json:"proc_id,omitempty"`   // 流程ID
	Starter  string         `json:"starter,omitempty"`   // 启动人人组
	BusiKey  string         `json:"busi_key,omitempty"`  // 业务单据ID
	BusiType string         `json:"busi_type,omitempty"` // 业务单据类型
	Input    map[string]any `json:"input,omitempty"`     // 启动参数 map[string]any

	Exec
	bpmn.TStartEvent
}

func (n StartProcCmd) check() error {
	if n.BusiKey == "" {
		return errors.New("参数BusiKey为空")
	}
	if n.BusiType == "" {
		return errors.New("参数BusiType为空")
	}
	if n.Starter == "" {
		return errors.New("参数Starter为空")
	}

	return nil
}

func (n *StartProcCmd) Bind(ctx context.Context, e *Engine) error {
	if err := n.check(); err != nil {
		return err
	}

	var t *Template
	if t = e.Templates.GetTemplate(n.ProcID); t == nil {
		return fmt.Errorf("未找到流程(ProcID = %s)", n.ProcID)
	}

	var ok bool
	if n.TStartEvent, ok = t.FindNormalStartEvent(); !ok {
		return errors.New("没有找到合适的StartEvnet")
	}

	n.Exec.Input = Merge(n.Exec.Input, n.Input)

	n.Exec.ProcInst = &ProcInst{
		InstID:   e.NewID(),
		ProcID:   n.ProcID,
		BusiKey:  n.BusiKey,
		BusiType: n.BusiType,
		Starter:  n.Starter,

		Init: maps.Clone(n.Exec.Input),

		Template: t,
		Engine:   e,
	}

	return nil
}

func (n *StartProcCmd) Do(ctx context.Context) error {
	n.Exec.ProcInst.StartTime = time.Now()
	n.Exec.ProcInst.Status = STATUS_START
	return n.CreateProcInst(ctx, n.ProcInst)
}

func (n *StartProcCmd) Emit(ctx context.Context, emit Emitter) error {
	return emit.Emit(fromOuter(ctx, n.Exec, n))
}

func (n StartProcCmd) Type() ActivityType {
	return OP_START_EVENT
}

type EndEventOp struct {
	Exec
	bpmn.TEndEvent

	UnimplementedActivity
}

func (n *EndEventOp) Do(ctx context.Context) error {
	n.ProcInst.Status = STATUS_FINISH
	n.ProcInst.EndTime = time.Now()

	return n.EndProcInst(ctx, n.ProcInst)
}
func (n EndEventOp) Type() ActivityType {
	return OP_END_EVENT
}

func doIntermediationThrowEvent(exec Exec, i bpmn.TIntermediateThrowEvent) Activity {
	if i.HasLinkEventDefinition() {
		return &linkEventOp{Exec: exec, Throw: i.GetLinkEventDefinition()}
	}
	panic("")
}

type linkEventOp struct {
	Exec
	Throw bpmn.TLinkEventDefinition

	UnimplementedActivity
}

func (n linkEventOp) Emit(ctx context.Context, emt Emitter) error {
	for _, c := range n.Template.Definitions.Process.IntermediateCatchEvent {
		if c.HasLinkEventDefinition() {
			if c.GetLinkEventDefinition().Name == n.Throw.Name {
				return emt.Emit(fromExec(n.Exec, c.OutgoingAssociation[0]))
			}
		}
	}
	panic("LinkEvent错误，Throw 没有找到 Catch")
}
