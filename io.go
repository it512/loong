package loong

import (
	"context"

	"github.com/it512/loong/bpmn/zeebe"
)

type TaskDefinition interface {
	Type() (string, error)
}

type IoOperator interface {
	BpmnElement
	GetTaskDefinition(context.Context) TaskDefinition
	GetInput() Getter
	GetTaskHeader(string) (string, bool)
	GetProperty(string) (string, bool)

	SetResult(Getter)
}

type IoConnector interface {
	Call(context.Context, IoOperator) error
}

type taskDef struct {
	typ string
	err error

	el   string
	eval ActivationEvaluator
	c    context.Context

	id string
}

func newTaskDef(ctx context.Context, eval ActivationEvaluator, el string) *taskDef {
	return &taskDef{
		el:   el,
		eval: eval,
		c:    ctx,
	}
}

func (t *taskDef) Type() (string, error) {
	if t.el == "" {
		return t.id, nil
	}

	if t.typ != "" {
		return t.typ, t.err
	}

	t.typ, _, t.err = eval[string](t.c, t.eval, t.el)
	return t.typ, t.err
}

type nopIo struct{}

func (nopIo) Call(_ context.Context, _ IoOperator) error { return nil }

type LazyGetFunc func(key string) (any, error)
type LazyBag struct {
	m map[string]any
}

func NewLazyBag() *LazyBag {
	return &LazyBag{
		m: make(map[string]any),
	}
}

func (b *LazyBag) Get(key string) (any, bool) {
	v, ok := b.m[key]
	if !ok {
		return nil, false
	}

	if f, ok := v.(LazyGetFunc); ok {
		a, err := f(key)
		if err != nil {
			panic(err)
		}
		return a, true
	}

	return v, true
}

func (b *LazyBag) Set(key string, val any) {
	b.m[key] = val
}

type InOut struct {
	in  *LazyBag
	out Getter
}

func newInOut() InOut {
	return InOut{
		in:  NewLazyBag(),
		out: emptyVar,
	}
}

func (io InOut) GetInput() Getter {
	return io.in
}

func (io *InOut) SetResult(result Getter) {
	io.out = result
}

func (io InOut) Get(key string) (any, bool) {
	k, _ := exp(key)
	return io.out.Get(k)
}

func (io *InOut) Set(key string, val any) {
	io.in.Set(key, val)
}

type ioer interface {
	GetIoInput() []zeebe.TIoMapping
	GetIoOutput() []zeebe.TIoMapping
	ActivationEvaluator
	IoConnector

	Getter
	Setter

	IoOperator
}

func lazy(ctx context.Context, eval ActivationEvaluator, el string) LazyGetFunc {
	return func(_ string) (any, error) {
		return eval.Eval(ctx, el)
	}
}

func in(ctx context.Context, in []zeebe.TIoMapping, eval ActivationEvaluator, s Setter) error {
	return Each(in, func(m zeebe.TIoMapping, _ int) error {
		if m.Target != "" {
			s.Set(m.Target, lazy(ctx, eval, m.Source))
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

func io(ctx context.Context, x ioer, s Setter) (err error) {
	if err = in(ctx, x.GetIoInput(), x, x); err != nil {
		return
	}
	if err = x.Call(ctx, x); err != nil {
		return
	}
	if err = out(x.GetIoOutput(), x, s); err != nil {
		return
	}
	return
}
