package loong

import (
	"context"

	"github.com/google/uuid"
)

type EventHandler interface {
	Handle(context.Context, Activity)
}

type IDGen interface {
	NewID() string
}

type uid struct{}

func (uid) NewID() string { return uuid.Must(uuid.NewV7()).String() }

type Emitter interface {
	Emit(...Activity) error
}

type Driver interface {
	Emitter
	Run() error
}

type Store interface {
	LoadProcInst(ctx context.Context, instID string, pi *ProcInst) error
	CreateProcInst(ctx context.Context, procInst *ProcInst) error
	EndProcInst(ctx context.Context, procInst *ProcInst) error

	CreateTasks(ctx context.Context, tasks ...UserTask) error
	LoadUserTask(ctx context.Context, taskID string, ut *UserTask) error
	EndUserTask(ctx context.Context, ut UserTask) error

	LoadUserTaskBatch(ctx context.Context, batchNO string) ([]UserTask, error)
	EndUserTaskBatch(ctx context.Context, batchNO string) error

	ForkExec(ctx context.Context, xs []Exec) error
	JoinExec(ctx context.Context, ex *Exec) error
	LoadForks(ctx context.Context, forkID string) ([]Exec, error)
	LoadExec(ctx context.Context, execID string, ex *Exec) error
}

func Each[S ~[]E, E any](s S, f func(E, int) error) error {
	for i, item := range s {
		if err := f(item, i); err != nil {
			return err
		}
	}
	return nil
}

func Must[T any](t T, err error) T {
	if err != nil {
		panic(err)
	}
	return t
}
