package loong

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/it512/loong/bpmn"
)

const (
	STATUS_START  = 1
	STATUS_FINISH = 100
)

type ActivityType string

const (
	NotApplicable  ActivityType = "N/A"
	OP_START_EVENT ActivityType = "OP_START_EVENT"
	OP_END_EVENT   ActivityType = "OP_END_EVENT"

	backgroundCmd ActivityType = "bgCmd" // for bgCmd only
)

type BpmnElement interface{ bpmn.BaseElement }

type Cmd interface {
	Init(context.Context, *Engine) error
	Do(context.Context) error
}

type Activity interface {
	Do(context.Context) error
	Emit(context.Context, Emitter) error
	Type() ActivityType
}

type ActivityCmd interface {
	Activity
	Init(context.Context, *Engine) error
}

type UnimplementedActivity struct{}

func (UnimplementedActivity) Do(ctx context.Context) error                { return nil }
func (UnimplementedActivity) Emit(ctx context.Context, emt Emitter) error { return nil }
func (UnimplementedActivity) Type() ActivityType                          { return NotApplicable }

type EventHandler interface {
	Handle(context.Context, Activity)
}
type eh struct{}

func (eh) Handle(ctx context.Context, op Activity) {}

type IDGen interface {
	NewID() string
}

type uid struct{}

func (uid) NewID() string { return uuid.Must(uuid.NewV7()).String() }

type Emitter interface {
	Emit(...Activity) error
}

type Backgrounder interface {
	Background(...Cmd) error
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
