package loong

import "context"

type InOut struct {
	in  *LazyBag
	out Var
}

func newInOut() InOut {
	return InOut{
		in:  NewLazyBag(),
		out: NewVar(),
	}
}

func (io InOut) GetInput(key string) (any, bool) {
	return io.in.Get(key)
}

func (io *InOut) SetResult(key string, val any) {
	io.out.Set(key, val)
}

func (io InOut) Get(key string) (any, bool) {
	k, _ := exp(key)
	return io.out.Get(k)
}

func (io *InOut) Set(key string, val any) {
	io.in.Set(key, val)
}

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

func lazy(ctx context.Context, eval ActivationEvaluator, el string) LazyGetFunc {
	return func(_ string) (any, error) {
		return eval.Eval(ctx, el)
	}
}
