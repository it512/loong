package loong

import (
	"context"
	"fmt"
	"strings"

	"github.com/expr-lang/expr"
	"github.com/expr-lang/expr/vm"
)

type Evaluator interface {
	Eval(ctx context.Context, el string, a any) (any, error)
}

type ActivationEvaluator interface {
	Eval(ctx context.Context, el string) (any, error)
}

func eval[T any](ctx context.Context, ae ActivationEvaluator, el string) (val T, a any, err error) {
	a, err = ae.Eval(ctx, el)
	if err != nil {
		return
	}

	var ok bool
	if val, ok = a.(T); !ok {
		err = fmt.Errorf("不能将 %T 转换为 %T", a, val)
	}
	return
}

func eval2[T any](ctx context.Context, e Evaluator, el string, env any) (val T, a any, err error) {
	a, err = e.Eval(ctx, el, env)
	if err != nil {
		return
	}

	var ok bool
	if val, ok = a.(T); !ok {
		err = fmt.Errorf("不能将 %T 转换为 %T", a, val)
	}
	return
}

type ExprEval struct {
	inst *vm.VM
}

func NewExprEval() *ExprEval {
	return &ExprEval{
		inst: &vm.VM{},
	}
}

func (e *ExprEval) Eval(ctx context.Context, ex string, a any) (any, error) {
	var el string
	var ok bool
	if el, ok = exp(ex); !ok {
		return el, nil
	}

	program, err := expr.Compile(el)
	if err != nil {
		return nil, err
	}

	return e.inst.Run(program, a)
}

func exp(s string) (string, bool) {
	a, ok1 := fx(s)
	b, ok2 := el(a)
	return b, ok1 || ok2
}

func fx(s string) (string, bool) {
	fx := strings.TrimSpace(s)
	return strings.CutPrefix(fx, "=")
}

func el(s string) (string, bool) {
	fx := strings.TrimSpace(s)
	if strings.HasPrefix(fx, "${") && strings.HasSuffix(fx, "}") {
		if a, ok := strings.CutPrefix(fx, "${"); ok {
			return strings.CutSuffix(a, "}")
		}
	}
	return fx, false
}
