package loong

import "context"

type Setter interface {
	Set(string, any)
}

type Getter interface {
	Get(k string) (v any, ok bool)
}

type Var map[string]any

func NewVar() Var {
	return make(Var)
}

func (b Var) Set(k string, v any) {
	b.Put(k, v)
}

func (b Var) Put(k string, v any) Var {
	b[k] = v
	return b
}

func (b Var) Get(k string) (v any, ok bool) {
	v, ok = b[k]
	return
}

func (b Var) Range(f func(string, any) error) error {
	for k, v := range b {
		if err := f(k, v); err != nil {
			return err
		}
	}
	return nil
}

func Merge(dest, src Var) Var {
	if dest == nil {
		dest = NewVar()
	}
	if src == nil {
		return dest
	}
	for k, v := range src {
		dest[k] = v
	}
	return dest
}

type Variable struct {
	Param Var
	Input Var
	Exec

	isChanged bool
}

func (v Variable) Changed() bool {
	return v.isChanged
}

func (v *Variable) PutVar(key string, val any) {
	v.isChanged = true
	v.Exec.ProcInst.Var.Set(key, val)
}

func (v Variable) Eval(ctx context.Context, el string) (any, error) {
	return v.Evaluator.Eval(ctx, el, v)
}
