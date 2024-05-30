package loong

import (
	"context"
	"maps"
	"time"

	"github.com/it512/loong/bpmn"
)

const (
	OP_START_EVENT ActivityType = "OP_START_EVENT"
	OP_END_EVENT   ActivityType = "OP_END_EVENT"
)

type StartEventOp struct {
	Exec
	bpmn.TStartEvent

	cmd StartProcCmd
}

func (n *StartEventOp) Do(ctx context.Context) error {
	n.ProcInst.InstID = n.Engine.NewID()
	n.ProcInst.ProcID = n.cmd.ProcID
	n.ProcInst.BusiKey = n.cmd.BusiKey
	n.ProcInst.BusiType = n.cmd.BusiType
	n.ProcInst.Starter = n.cmd.Starter

	n.ProcInst.Template = n.GetTemplate(n.ProcID)

	var ok bool
	if n.TStartEvent, ok = n.ProcInst.Template.FindNormalStartEvent(); !ok {
		panic("没有找到合适的StartEvnet")
	}

	n.Exec.Input = Merge(n.Exec.Input, n.cmd.Input)

	n.ProcInst.Init = maps.Clone(n.Exec.Input)
	n.ProcInst.StartTime = time.Now()
	n.ProcInst.Status = STATUS_START

	n.Exec.Status = STATUS_START

	return n.CreateProcInst(ctx, n.ProcInst)
}

func (n *StartEventOp) Emit(ctx context.Context, emit Emitter) error {
	return emit.Emit(fromOuter(ctx, n.Exec, n))
}

func (n StartEventOp) Type() ActivityType {
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

	// TODO
	/*
		if n.TEndEvent.HasErrorEventDefinition() {
			if e, ok := n.Template.FindError(n.TEndEvent.GetErrorEventDefinition().ErrorRef); ok {
				n.ProcInst.EndCode = e.ErrorCode
				n.ProcInst.EndName = e.Name
				n.EndType = "error"
			}
		}
	*/

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
