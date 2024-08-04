package loong

import (
	"testing"
)

func TestVarChange(t *testing.T) {
	var v Variable
	v.Exec.ProcInst = &ProcInst{Var: NewVar()}
	v.Put("Var.x", "x")
	if !v.isVarChanged {
		t.Errorf("not changed")
	}
}
