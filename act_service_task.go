package loong

import (
	"context"

	"github.com/it512/loong/bpmn"
	"github.com/it512/loong/bpmn/zeebe"
)

type ServiceTask struct {
	Exec
	bpmn.TServiceTask

	in  Var
	out Getter
}

func (s *ServiceTask) Do(ctx context.Context) error {
	s.in = NewVar()
	s.out = emptyVar

	if err := in(ctx, s.TServiceTask.Input, s.Exec, s.in); err != nil {
		return nil
	}

	if err := s.connector.Call(ctx, s); err != nil {
		return err
	}

	return out(s.TServiceTask.Output, s.out, s.Exec.Input)

}

func (s *ServiceTask) Emit(ctx context.Context, emt Emitter) error {
	return s.EmitDefault(ctx, s.TServiceTask, emt)
}

func (s *ServiceTask) GetInput() Getter {
	return s.in
}

func (s *ServiceTask) SetResult(out Getter) {
	s.out = out
}

func (s *ServiceTask) GetTaskDefinition(ctx context.Context) TaskDefinition {
	return newTaskDef(
		ctx,
		s,
		s.TServiceTask.TaskDefinition.TypeName,
	)
}

func (s ServiceTask) GetTaskHeader(key string) (string, bool) {
	return zeebe.GetTaskHeader(s.TaskHeaders, key)
}

func (s ServiceTask) GetProperty(name string) (string, bool) {
	return zeebe.GetProperty(s.Properties, name)
}
