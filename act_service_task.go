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
	if err := io(ctx, s, s); err != nil {
		return err
	}

	if s.Variable.Changed() {
		if err := s.Storer.SaveVar(ctx, s.ProcInst); err != nil {
			return err
		}
	}

	/*
		if e, ok := err.(*BizErr); ok {
			for _, t := range s.Exec.ProcInst.Template.Definitions.Errors {
				if cmp.Compare(t.ErrorCode, e.ErrorCode) == 0 {
					return &actErrOp{
						error: e,
						Exec:  s.Exec,
					}
				}
			}
			return &actErrOp{
				error: e,
				Exec:  s.Exec,
			}
		}
	*/

	return nil
}

func (s *serviceTaskOp) Emit(ctx context.Context, emt Emitter) error {
	return emt.Emit(fromOuter(ctx, s.Exec, s))
}

func (s *serviceTaskOp) GetTaskDefinition(ctx context.Context) TaskDefinition {
	return newTaskDef(
		ctx,
		s,
		s.TServiceTask.TaskDefinition,
	)
}
