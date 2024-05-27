package loong

import (
	"context"
	"fmt"
	"sync"

	"github.com/it512/loong/bpmn"
)

func choose(ctx context.Context, ae ActivationEvaluator, in []bpmn.TSequenceFlow) (out []bpmn.TSequenceFlow, err error) {
	if len(in) == 1 {
		return in, nil
	}

	for _, flow := range in {
		if flow.HasConditionExpression() {
			var b bool
			if b, _, err = eval[bool](ctx, ae, flow.GetConditionExpression()); err != nil {
				return
			}
			if b {
				out = append(out, flow)
			}
		}
	}
	if len(out) == 0 {
		out = in
	}
	return
}

func chooseDefault(me any, flows []bpmn.TSequenceFlow) bpmn.TSequenceFlow {
	if len(flows) == 1 {
		return flows[0]
	}
	if d, ok := me.(bpmn.DefaultAttrElement); ok {
		if key := d.GetDefault(); key != "" {
			if ele, ok := bpmn.Find(flows, key); ok {
				return ele
			}
		}
	}
	for _, flow := range flows {
		if !flow.HasConditionExpression() {
			return flow
		}
	}
	return flows[0] // 默认返回第一条
}

var xPool = sync.Pool{
	New: func() any {
		return &sequenceFlow{}
	},
}

type sequenceFlow struct {
	Exec
	bpmn.TSequenceFlow

	target BpmnElement
}

func (c *sequenceFlow) Do(_ context.Context) error {
	c.InTag = c.GetId()

	var ok bool
	c.target, ok = c.ProcInst.Template.FindElementByID(c.TargetRef)
	if !ok {
		panic(fmt.Errorf("未找到目标 TargetRef = %s", c.TargetRef))
	}

	c.Exec.elementID = c.target.GetId()

	return nil
}

func (c *sequenceFlow) Emit(_ context.Context, commit Emitter) (err error) {
	switch c.target.GetType() {
	case bpmn.UserTask:
		err = commit.Emit(&userTaskOp{UserTask: UserTask{Exec: c.Exec}, InOut: newInOut(), TUserTask: bpmn.Cast[bpmn.TUserTask](c.target)})
	case bpmn.ExclusiveGateway:
		err = commit.Emit(&exclusivGatewayOp{gateway: gateway{Forker: bpmn.Cast[bpmn.TExclusiveGateway](c.target), Exec: c.Exec}})
	case bpmn.ParallelGateway:
		err = commit.Emit(&parallelGatewayCmd{gateway: gateway{Forker: bpmn.Cast[bpmn.TParallelGateway](c.target), Exec: c.Exec}})
	case bpmn.ServiceTask:
		err = commit.Emit(&ServiceTask{Exec: c.Exec, InOut: newInOut(), TServiceTask: bpmn.Cast[bpmn.TServiceTask](c.target)})
	case bpmn.EndEvent:
		err = commit.Emit(&EndEventOp{Exec: c.Exec, TEndEvent: bpmn.Cast[bpmn.TEndEvent](c.target)})
	case bpmn.IntermediateThrowEvent:
		op := doIntermediationThrowEvent(c.Exec, bpmn.Cast[bpmn.TIntermediateThrowEvent](c.target))
		err = commit.Emit(op)
	case bpmn.Task:
		err = commit.Emit(&taskOp{Exec: c.Exec, Ele: bpmn.Cast[bpmn.TTask](c.target)})
	default:
		panic(fmt.Errorf("不支持的类型 Type: %s, ID: %s", c.target.GetType(), c.target.GetId()))
	}

	xPool.Put(c)
	return
}

func fromExec(ex Exec, out string) *sequenceFlow {
	if f, ok := ex.Template.FindSequenceFlow(out); ok {
		sf := xPool.Get().(*sequenceFlow)
		sf.Exec = ex
		sf.TSequenceFlow = f
		return sf
		// return &sequenceFlow{Exec: ex, TSequenceFlow: f}
	}
	panic("未找到Sequenceflow")
}
