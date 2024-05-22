package io

import (
	"context"

	"github.com/it512/loong"
)

type Io struct {
}

func (Io) Call(ctx context.Context, o loong.IoOperator) error {
	v := loong.NewVar()
	v.Put("xx", []string{"a", "b"})
	o.SetResult(v)
	return nil
}
