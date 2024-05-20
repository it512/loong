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

type Evaluator interface {
	Eval(ctx context.Context, el string, a any) (any, error)
}

type ActivationEvaluator interface {
	Eval(ctx context.Context, el string) (any, error)
}

func eval[T any](ctx context.Context, ae ActivationEvaluator, el string) (val T, a any, err error) {
	a, err = ae.Eval(ctx, el)
	if err != nil {
		return
	}

	val = a.(T)
	return
}

func eval2[T any](ctx context.Context, e Evaluator, el string, env any) (val T, a any, err error) {
	a, err = e.Eval(ctx, el, env)
	if err != nil {
		return
	}

	val = a.(T)
	return
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
