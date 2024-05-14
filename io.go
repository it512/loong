package loong

import (
	"context"

	"github.com/it512/loong/bpmn/zeebe"
)

type TaskDefinition interface {
	Type() (string, error)
}

type Operator interface {
	BpmnElement
	GetTaskDefinition(context.Context) TaskDefinition
	GetInput() Getter
	GetTaskHeader(string) (string, bool)
	GetProperty(string) (string, bool)

	SetResult(Getter)
}

type ServiceConnector interface {
	Call(context.Context, Operator) error
}

type emptyTaskDef struct {
	id string
}

func (t emptyTaskDef) Type() (string, error) {
	return t.id, nil
}

type taskDef struct {
	typ string
	err error

	el   string
	eval ActivationEvaluator
	c    context.Context
}

func newTaskDef(ctx context.Context, eval ActivationEvaluator, el string) *taskDef {
	return &taskDef{
		el:   el,
		eval: eval,
		c:    ctx,
	}
}

func (t *taskDef) Type() (string, error) {
	if t.typ != "" {
		return t.typ, t.err
	}

	a, err := t.eval.Eval(t.c, t.el)
	t.err = err
	if err != nil {
		return "", err
	}
	t.typ = a.(string)
	return t.typ, t.err
}

type nilConnect struct{}

func (nilConnect) Call(_ context.Context, _ Operator) error { return nil }

var emptyConnect = new(nilConnect)

func in(ctx context.Context, in []zeebe.TIoMapping, eval ActivationEvaluator, s Setter) error {
	return Each(in, func(m zeebe.TIoMapping, _ int) error {
		if m.Target != "" {
			a, err := eval.Eval(ctx, m.Source)
			if err != nil {
				return err
			}
			s.Set(m.Target, a)
		}
		return nil
	})

}

func out(o []zeebe.TIoMapping, g Getter, s Setter) error {
	return Each(o, func(m zeebe.TIoMapping, _ int) error {
		if m.Target != "" {
			if v, ok := g.Get(m.Source); ok {
				s.Set(m.Target, v)
			}
		}
		return nil
	})
}
