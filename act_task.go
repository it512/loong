package loong

import (
	"context"

	"github.com/it512/loong/bpmn"
)

type taskOp struct {
	Exec
	Ele bpmn.TTask
}

func (t taskOp) Emit(ctx context.Context, emt Emitter) error {
	return t.EmitDefault(ctx, t.Ele, emt)
}
