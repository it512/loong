package loong

import (
	"context"
	"fmt"
	"strings"
)

type Setter interface {
	Set(string, any)
}

type Getter interface {
	Get(k string) (v any, ok bool)
}

type Putter interface {
	Put(string, any)
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

func Merge(dest, src Var) Var {
	if dest == nil {
		dest = NewVar()
	}
	if len(src) == 0 {
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

	BpmnElement

	isVarChanged bool
}

func (v *Variable) Put(key string, val any) {
	part := strings.Split(key, ".")
	if len(part) < 2 {
		panic("设置变量时必须指定作用域")
	}

	v.PutScope(part[0], part[1], val)
}

func (v *Variable) PutScope(s, k string, val any) {
	switch s {
	case "Var":
		v.PutVar(k, val)
	case "Input":
		v.PutInput(k, val)
	default:
		panic(fmt.Errorf("作用域: %s 不支持写入", s))
	}
}

func (v *Variable) PutVar(key string, val any) {
	v.isVarChanged = true
	v.Exec.ProcInst.Var.Set(key, val)
}

func (v *Variable) PutInput(key string, val any) {
	v.Input.Set(key, val)
}

func (v *Variable) PutParam(key string, val any) {
	v.Param.Set(key, val)
}

func (v Variable) Eval(ctx context.Context, el string) (any, error) {
	return v.Evaluator.Eval(ctx, el, v)
}

func (v Variable) saveTo(ctx context.Context, s Storer) error {
	if v.isVarChanged {
		return s.SaveVar(ctx, v.ProcInst)
	}
	return nil
}
