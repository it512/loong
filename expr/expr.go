package expr

import (
	"context"
	"strings"

	"github.com/expr-lang/expr"
	"github.com/expr-lang/expr/vm"
)

type ExprEval struct {
	//prgMap map[string]*vm.Program
	inst *vm.VM
}

func New() *ExprEval {
	return &ExprEval{
		inst: &vm.VM{},
		//prgMap: make(map[string]*vm.Program),
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

func el(s string) (string, bool) {
	fx := strings.TrimSpace(s)
	if strings.HasPrefix(fx, "${") && strings.HasSuffix(fx, "}") {
		if a, ok := strings.CutPrefix(fx, "${"); ok {
			return strings.CutSuffix(a, "}")
		}
	}
	return fx, false
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
