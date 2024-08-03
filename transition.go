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

func chooseDefault(me outer, flows []bpmn.TSequenceFlow) bpmn.TSequenceFlow {
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

type sequenceFlow struct {
	Variable
	bpmn.TSequenceFlow
	target BpmnElement

	UnimplementedActivity
}

func (c *sequenceFlow) Do(_ context.Context) error {
	c.Exec.InTag = c.GetId()

	var ok bool
	c.target, ok = c.ProcInst.Template.FindElementByID(c.TargetRef)
	if !ok {
		panic(fmt.Errorf("未找到目标 TargetRef = %s", c.TargetRef))
	}

	c.elementID = c.target.GetId()
	c.elementType = c.target.GetType()

	return nil
}

func (c *sequenceFlow) Emit(_ context.Context, commit Emitter) (err error) {
	switch c.target.GetType() {
	case bpmn.UserTask:
		err = commit.Emit(&userTaskOp{UserTask: UserTask{Variable: c.Variable}, InOut: newInOut(), TUserTask: bpmn.Cast[bpmn.TUserTask](c.target)})
	case bpmn.ExclusiveGateway:
		err = commit.Emit(&exclusivGatewayOp{Variable: c.Variable, TExclusiveGateway: bpmn.Cast[bpmn.TExclusiveGateway](c.target)})
	case bpmn.ParallelGateway:
		err = commit.Emit(&parallelGatewayCmd{TParallelGateway: bpmn.Cast[bpmn.TParallelGateway](c.target), Exec: c.Exec})
	case bpmn.ServiceTask:
		err = commit.Emit(&serviceTaskOp{Variable: c.Variable, InOut: newInOut(), TServiceTask: bpmn.Cast[bpmn.TServiceTask](c.target)})
	case bpmn.EndEvent:
		err = commit.Emit(&EndEventOp{Exec: c.Exec, TEndEvent: bpmn.Cast[bpmn.TEndEvent](c.target)})
	case bpmn.IntermediateThrowEvent:
		op := doIntermediationThrowEvent(c.Variable, bpmn.Cast[bpmn.TIntermediateThrowEvent](c.target))
		err = commit.Emit(op)
	case bpmn.Task:
		err = commit.Emit(&taskOp{Variable: c.Variable, TTask: bpmn.Cast[bpmn.TTask](c.target)})
	default:
		panic(fmt.Errorf("不支持的类型 Type: %s, ID: %s", c.target.GetType(), c.target.GetId()))
	}

	putToPool(c)
	return
}

var sfPool = sync.Pool{
	New: func() any {
		return &sequenceFlow{}
	},
}

func getFromPool(v Variable, f bpmn.TSequenceFlow) *sequenceFlow {
	sf := sfPool.Get().(*sequenceFlow)

	sf.Variable.Exec = v.Exec
	sf.Variable.Input = v.Input
	sf.Variable.isChanged = false

	sf.TSequenceFlow = f

	return sf
}

func putToPool(sf *sequenceFlow) {
	sfPool.Put(sf)
}

type outer interface {
	GetOutgoingAssociation() []string
	FindSequenceFlow(string) (bpmn.TSequenceFlow, bool)
	FindSequenceFlows([]string) []bpmn.TSequenceFlow
	ActivationEvaluator
}

/*
	func fromExec(ex Exec, out string) *sequenceFlow {
		if f, ok := ex.Template.FindSequenceFlow(out); ok {
			return getFromPool(ex, f)
		}
		panic("未找到Sequenceflow")
	}
*/

func fromVariable(v Variable, out string) *sequenceFlow {
	if f, ok := v.ProcInst.Template.FindSequenceFlow(out); ok {
		return getFromPool(v, f)
	}
	panic("未找到Sequenceflow")
}

func fromOuter(ctx context.Context, v Variable, o outer) *sequenceFlow {
	flows := o.FindSequenceFlows(o.GetOutgoingAssociation())
	out, err := choose(ctx, o, flows)
	if err != nil {
		panic(err)
	}
	f := chooseDefault(o, out)
	return getFromPool(v, f)
}
