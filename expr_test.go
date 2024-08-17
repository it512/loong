package loong

import (
	"testing"
)

func TestEL(t *testing.T) {
	el := "=a"
	s, ok := Expr(el)
	if !ok {
		t.Error("xxxx")
	}

	if s != "a" {
		t.Error("yyyy")
	}
}

func TestEmptyEL(t *testing.T) {
	el := "= "
	s, ok := Expr(el)
	if ok {
		t.Error("xxxx")
	}

	if s != "" {
		t.Error("yyyy")
	}
}
