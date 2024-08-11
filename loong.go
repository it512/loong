package loong

import (
	"cmp"
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/it512/loong/bpmn"
)

const (
	// 系统内置的账号，用于自动化任务
	Robot = "robot"
)

const (
	STATUS_START  = 1
	STATUS_FINISH = 100
)

type ActivityType string

func (at ActivityType) Type() (string, error) {
	return string(at), nil
}

func (at ActivityType) String() string {
	return string(at)
}

func (at ActivityType) Eq(a any) bool {
	switch v := a.(type) {
	case string:
		return cmp.Compare(at.String(), v) == 0
	case ActivityType:
		return at == v
	}
	return false
}

const (
	AT_NotApplicable ActivityType = "N/A"
	OP_START_EVENT   ActivityType = "OP_START_EVENT"
	OP_END_EVENT     ActivityType = "OP_END_EVENT"

	AT_USER_TASK_COMMIT ActivityType = "USER_TASK_COMMIT"
	AT_USER_TASK        ActivityType = "USER_TASK"
)

type BpmnElement interface{ bpmn.BaseElement }

type Cmd interface {
	Bind(context.Context, *Engine) error
	Do(context.Context) error
}

type Activity interface {
	Do(context.Context) error
	Emit(context.Context, Emitter) error
	Type() ActivityType
}

type ActivityCmd interface {
	Activity
	Bind(context.Context, *Engine) error
}

type UnimplementedActivity struct{}

func (UnimplementedActivity) Do(ctx context.Context) error                       { return nil }
func (UnimplementedActivity) Emit(ctx context.Context, emt Emitter) error        { return nil }
func (UnimplementedActivity) Type() ActivityType                                 { return AT_NotApplicable }
func (u UnimplementedActivity) GetTaskDefinition(context.Context) TaskDefinition { return u.Type() }

type EventHandler interface {
	Handle(context.Context, Activity)
}
type eh struct{}

func (eh) Handle(ctx context.Context, op Activity) {}

type IDGenerator interface {
	NewID() string
}

type uid struct{}

func (uid) NewID() string { return uuid.Must(uuid.NewV7()).String() }

type Emitter interface {
	Emit(...Activity) error
}

type Storer interface {
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
	LoadForks(ctx context.Context, instID, forkID string) ([]Exec, error)
	LoadExec(ctx context.Context, instID, execID string, ex *Exec) error

	SaveVar(ctx context.Context, procInst *ProcInst) error
}

type Txer interface {
	DoTx(context.Context, func(context.Context) error) error
}

type Xx interface {
	Do(context.Context, Exec, BpmnElement) (string, string, error)
}

type X struct{}

func (X) Do(ctx context.Context, v Variable) (owner, manager string, err error) {
	return v.Starter, v.Starter, nil
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

const (
	banner = `
   __
  / /  _   _   _     _
 / /_,'o|,'o| / \/7,'o|
/___/|_,'|_,'/_n_/ |_,'
                   _//  (2024)
BPMN2流程引擎(%s)
-------------------------------
`
	Version = "0.0.0"
)

func logo() {
	log.Printf(banner, "v"+Version)
}
