package loong

import (
	"context"
	"testing"
)

func TestEmptyMap(t *testing.T) {
	eval := NewExprEval()
	var v Variable
	// v.Input = Var{"op": "op"}
	_, err := eval.Eval(context.Background(), "${Input.op}", v)
	if err != nil {
		t.Error(err)
	}
}
