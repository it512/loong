package loong

import (
	"context"

	"github.com/it512/loong/bpmn"
)

type ServiceTask struct {
	Exec
	bpmn.TServiceTask

	InOut
}

func (s *ServiceTask) Do(ctx context.Context) error {
	return io(ctx, s, s.Exec.Input)
}

func (s *ServiceTask) Emit(ctx context.Context, emt Emitter) error {
	return emt.Emit(fromOuter(ctx, s.Exec, s))
}

func (s *ServiceTask) GetTaskDefinition(ctx context.Context) TaskDefinition {
	return newTaskDef(
		ctx,
		s,
		s.TServiceTask.TaskDefinition.TypeName,
	)
}
