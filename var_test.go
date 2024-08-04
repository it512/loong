package loong

import (
	"testing"
)

func TestVarChange(t *testing.T) {
	var v Variable
	v.Exec.ProcInst = &ProcInst{Var: NewVar()}
	v.Put("x", "x")
	if !v.Changed() {
		t.Errorf("not changed")
	}
}
