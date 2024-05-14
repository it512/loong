package bpmn

import (
	"cmp"
)

type BaseElement interface {
	GetId() string
	GetName() string
	GetType() ElementType
}

type IDer interface {
	GetId() string
}

func Find[S ~[]E, E IDer](x S, target string) (ele E, ok bool) {
	for i, t := range x {
		if cmp.Compare(t.GetId(), target) == 0 {
			ele = x[i]
			ok = true
			return
		}
	}
	return
}

func FindSequenceFlows(definitions *TDefinitions, ids []string) (ret []TSequenceFlow) {
	for _, flow := range definitions.Process.SequenceFlows {
		for _, id := range ids {
			if id == flow.Id {
				ret = append(ret, flow)
			}
		}
	}
	return
}

func FindNormalStartEvent(definitions *TDefinitions) (TStartEvent, bool) {
	for _, s := range definitions.Process.StartEvents {
		if s.IsNormal() {
			return s, true
		}
	}
	return definitions.Process.StartEvents[0], false
}

type BpmnElement interface {
	TStartEvent |
		TUserTask | TServiceTask | TTask |
		TParallelGateway | TExclusiveGateway | TEventBasedGateway |
		TIntermediateThrowEvent | TIntermediateCatchEvent |
		TEndEvent
}

func Cast[E BpmnElement](b BaseElement) E {
	return b.(E)
}

func FindElementById(definitions *TDefinitions, id string) (BaseElement, bool) {
	if x, ok := Find(definitions.Process.UserTasks, id); ok {
		return x, ok
	}
	if x, ok := Find(definitions.Process.ServiceTasks, id); ok {
		return x, ok
	}
	if x, ok := Find(definitions.Process.Tasks, id); ok {
		return x, ok
	}
	if x, ok := Find(definitions.Process.ExclusiveGateway, id); ok {
		return x, ok
	}
	if x, ok := Find(definitions.Process.ParallelGateway, id); ok {
		return x, ok
	}
	if x, ok := Find(definitions.Process.EventBasedGateway, id); ok {
		return x, ok
	}
	if x, ok := Find(definitions.Process.IntermediateCatchEvent, id); ok {
		return x, ok
	}
	if x, ok := Find(definitions.Process.IntermediateTrowEvent, id); ok {
		return x, ok
	}
	if x, ok := Find(definitions.Process.EndEvents, id); ok {
		return x, ok
	}
	return definitions.Process.StartEvents[0], false
}
