package loong

import (
	"context"

	"github.com/it512/loong/bpmn/zeebe"
)

type IoSet []IoCaller

func (io IoSet) Call(ctx context.Context, task IoTasker) error {
	for _, call := range io {
		if err := call.Call(ctx, task); err != nil {
			return err
		}
	}
	return nil
}

type ioer interface {
	GetIoInput() []zeebe.TIoMapping
	GetIoOutput() []zeebe.TIoMapping
	ActivationEvaluator
	IoCaller

	Getter
	Setter

	IoTasker
}

func in(ctx context.Context, in []zeebe.TIoMapping, eval ActivationEvaluator, s Setter) error {
	return Each(in, func(m zeebe.TIoMapping, _ int) error {
		if m.Target != "" {
			s.Set(m.Target, lazy(ctx, eval, m.Source))
		}
		return nil
	})
}

func out(o []zeebe.TIoMapping, g Getter, s Putter) error {
	return Each(o, func(m zeebe.TIoMapping, _ int) error {
		if m.Target != "" {
			if v, ok := g.Get(m.Source); ok {
				s.Put(m.Target, v)
			}
		}
		return nil
	})
}

func io(ctx context.Context, x ioer, s Putter) (err error) {
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
