package loong

import (
	"context"

	"github.com/it512/loong/bpmn"
)

type serviceTaskOp struct {
	InOut
	Variable
	bpmn.TServiceTask

	UnimplementedActivity
}

func (s *serviceTaskOp) Do(ctx context.Context) error {
	if err := io(ctx, s, s); err != nil { // err == nil
		return w(err, s.GetId(), s.Variable)
	}
	if err := s.Storer.SaveVar(ctx, s.ProcInst); err != nil {
		return err
	}
	return nil
}

func (s *serviceTaskOp) Emit(ctx context.Context, emt Emitter) error {
	return emt.Emit(fromOuter(ctx, s.Variable, s))
}

func (s *serviceTaskOp) GetTaskDefinition(ctx context.Context) TaskDefinition {
	return newTaskDef(
		ctx,
		s,
		s.TServiceTask.TaskDefinition,
	)
}
