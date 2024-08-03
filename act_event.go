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
	Var      map[string]any `json:"var,omitempty"`       // 启动参数 map[string]any
	Tags     map[string]any `json:"tags,omitempty"`      // 流程标志 map[string]any

	Variable
	bpmn.TStartEvent
	InOut
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
	if t = e.GetTemplate(n.ProcID); t == nil {
		return fmt.Errorf("未找到流程(ProcID = %s)", n.ProcID)
	}

	if !t.Definitions.Process.IsExecutable {
		return fmt.Errorf("流程已停用(ProcID = %s)", n.ProcID)
	}

	var ok bool
	if n.TStartEvent, ok = t.FindNormalStartEvent(); !ok {
		return errors.New("没有找到合适的StartEvnet")
	}

	n.Variable.Input = maps.Clone(n.Input)
	n.Variable.Exec.ProcInst = &ProcInst{
		InstID:   e.NewID(),
		ProcID:   n.ProcID,
		BusiKey:  n.BusiKey,
		BusiType: n.BusiType,
		Starter:  n.Starter,

		Init: maps.Clone(n.Var),
		Var:  maps.Clone(n.Var),

		Tags: maps.Clone(n.Tags),

		Template: t,
		Engine:   e,
	}

	return nil
}

func (n *StartProcCmd) Do(ctx context.Context) error {
	// n.IoConnector.Call(ctx, n)
	n.Exec.ProcInst.StartTime = time.Now()
	n.Exec.ProcInst.Status = STATUS_START
	return n.Storer.CreateProcInst(ctx, n.Exec.ProcInst)
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

	if i.HasMessageEventDefinitio() {
		return &messageIntermediateThrowEventOp{Variable: Variable{Exec: exec}, InOut: newInOut(), TIntermediateThrowEvent: i}
	}
	panic("不支持的IntermediateThrowEvent类型")
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

type messageIntermediateThrowEventOp struct {
	bpmn.TIntermediateThrowEvent
	Variable
	InOut

	UnimplementedActivity
}

func (s *messageIntermediateThrowEventOp) Do(ctx context.Context) (err error) {
	if err = io(ctx, s, s); err != nil {
		return
	}

	if s.Variable.Changed() {
		if err = s.Storer.SaveVar(ctx, s.ProcInst); err != nil {
			return
		}
	}
	return
}

func (s *messageIntermediateThrowEventOp) Emit(ctx context.Context, emt Emitter) error {
	return emt.Emit(fromOuter(ctx, s.Exec, s))
}

func (s *messageIntermediateThrowEventOp) GetTaskDefinition(ctx context.Context) TaskDefinition {
	return newTaskDef(
		ctx,
		s,
		s.TaskDefinition,
	)
}
