package todo

import (
	"context"

	"github.com/it512/loong"
)

type ReassignOp struct {
	TaskID string
}

func (r ReassignOp) Do(ctx context.Context) error {
	return nil
}

func (c ReassignOp) Emit(ctx context.Context, emt loong.Emitter) error {
	return nil
}
