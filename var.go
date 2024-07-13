package loong

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
