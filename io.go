package loong

import (
	"context"

	"github.com/it512/loong/bpmn/zeebe"
)

type TaskDefinition interface {
	Type() (string, error)
}

type IoTasker interface {
	BpmnElement

	GetTaskDefinition(context.Context) TaskDefinition
	GetTaskHeader(string) (string, bool)
	GetProperty(string) (string, bool)

	GetInput(string) (any, bool)
	SetResult(string, any)
}

type IoCaller interface {
	Call(context.Context, IoTasker) error
}

type taskDef struct {
	typ string
	err error

	el   string
	eval ActivationEvaluator
	c    context.Context
}

func newTaskDef(ctx context.Context, eval ActivationEvaluator, td zeebe.TTaskDefinition) *taskDef {
	return &taskDef{
		el:   td.TypeName,
		eval: eval,
		c:    ctx,
	}
}

func (t *taskDef) Type() (string, error) {
	if t.typ != "" {
		return t.typ, t.err
	}

	t.typ, _, t.err = eval[string](t.c, t.eval, t.el)
	return t.typ, t.err
}
