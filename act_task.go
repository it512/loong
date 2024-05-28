package loong

import (
	"context"

	"github.com/it512/loong/bpmn"
)

type taskOp struct {
	Exec
	bpmn.TTask
}

func (t *taskOp) Emit(ctx context.Context, emt Emitter) error {
	return emt.Emit(fromOuter(ctx, t.Exec, t))
}
