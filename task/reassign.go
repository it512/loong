package task

import (
	"context"

	"github.com/it512/loong"
)

type ReassignOp struct {
	TaskID string

	loong.UnimplementedActivity
}

func (r ReassignOp) Do(ctx context.Context) error {
	return nil
}
